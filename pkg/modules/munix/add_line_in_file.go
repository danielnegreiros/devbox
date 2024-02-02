package munix

import (
	"io/fs"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/errs"
	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/ports"
)

// Take as parameters. path, regex, line, create, state, owner, group, mode, create, vaidate
// check if file exists
// if not and create == false, return error
// else
// if file exist, save its mode permissions
// read file content
// perform regex
// if found, regex replace it
// if not found append to last line
// if file does not exist, create it on temp mode
// save content to tmp file
// if validate diferent than null, validate it
// copy to permanent location wth proper permissions

type LineInFile struct {
	path         string
	regex        string
	line         string
	state        string
	owner        string
	group        string
	mode         string
	validate     string
	shouldCreate bool
	result       map[string]string
}

func NewLineInFile(path, regex, line, state, owner, group, mode, validate string, shouldCreate bool) *LineInFile {
	return &LineInFile{
		path:         path,
		regex:        regex,
		line:         line,
		state:        state,
		owner:        owner,
		group:        group,
		mode:         mode,
		validate:     validate,
		shouldCreate: shouldCreate,
		result:       map[string]string{},
	}
}

func (l *LineInFile) Execute(infra ports.InfraRepo) map[string]string {

	fileInfo, err := isFilePresent(l.path, infra)
	modes := strings.Split(strings.Trim(fileInfo, "\n"), " ")

	if err != nil && !l.shouldCreate {
		l.log("false", "file not present and create disabled", "")
	}

	if err != nil && l.shouldCreate {
		l.createIfRequested(modes[0], modes[1], modes[2])
	}

	l.regexIfFileExist(infra, modes[2])

	return l.result
}

func (l *LineInFile) log(success string, reason string, content string) {
	if success != "" {
		l.result["success"] = success
	}

	if reason != "" {
		l.result["reason"] = reason
	}

	if content != "" {
		l.result["content"] = content
	}

}

func (l *LineInFile) createIfRequested(user string, group string, mode string) {

	file, err := createFile(l.path)
	if err != nil {
		l.log("false", err.Error(), "")
	}
	defer file.Close()
	file.WriteString(l.line + "\n")
	l.log("true", "", l.line)

}

func (l *LineInFile) regexIfFileExist(infra ports.InfraRepo, mode string) {

	content, err := readFile(l.path, infra)
	if err != nil {
		l.log("false", err.Error(), "")
		return
	}

	regexExp := regexp.MustCompilePOSIX(l.regex)
	patExists := doesPatternExist(content, regexExp)

	switch patExists {
	case true:
		content = replaceRegex(content, regexExp, l.line)
	case false:
		content = appendToFile(content, l.line)
	}

	if l.validate != "" {
		err = validateFile(l.validate)
		errs.PanicIfErr(err)
	}

	fileMode, _ := strconv.Atoi(mode)
	err = writeFile(infra, l.path, content, fs.FileMode(fileMode))
	if err != nil {
		l.log("false", err.Error(), "")
	} else {
		l.log("true", "", content)
	}

}

func isFilePresent(path string, infra ports.InfraRepo) (string, error) {
	output := infra.ExecuteCommand("stat -c '%u %g %a' " + path)
	if output.Error != nil {
		errs.PanicIfErr(output.Error)
	}

	if output.Error != nil {
		return "", output.Error
	}

	message := strings.Trim(output.OutMessage, "\n")
	message = strings.Trim(message, "'")

	return message, nil
}

func readFile(path string, infra ports.InfraRepo) (string, error) {
	output := infra.ExecuteCommand("cat " + path)
	if output.Error != nil {
		return "", output.Error
	}

	return output.OutMessage, nil
}

func doesPatternExist(content string, pattern *regexp.Regexp) bool {
	return pattern.MatchString(content)
}

func replaceRegex(content string, regex *regexp.Regexp, line string) string {
	return regex.ReplaceAllString(content, line)
}

func appendToFile(content string, line string) string {
	if strings.HasSuffix(content, "\n") {
		return content + line + "\n"
	} else {
		return content + "\n" + line + "\n"
	}
}

func createFile(filepath string) (*os.File, error) {
	return os.Create(filepath)
}

func writeFile(infra ports.InfraRepo, filePath, content string, mode fs.FileMode) error {
	cmd := "echo \"" + strings.Trim(content, "\n") + "\" > " + filePath
	output := infra.ExecuteCommand(cmd)
	if output.Error != nil {
		return output.Error
	}
	return nil
}

func validateFile(command string) error {
	return nil
}
