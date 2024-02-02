package conf

type HardenConf struct {
	Files []File
}

type File struct {
	Path    string
	Changes []FileChange
}

type FileChange struct {
	Regex        string
	Line         string
	State        string
	Owner        string
	Group        string
	Mode         string
	Validate     string
	ShouldCreate bool
}

func GetDefaultHarden() HardenConf {

	// changes recommendations froom: https://www.digitalocean.com/community/tutorials/how-to-harden-openssh-on-ubuntu-20-04
	// book: Mastering Linux Security and Hardening
	sshdChanges := []FileChange{
		{
			Regex:        "^Protocol.*",
			Line:         "Protocol 2",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^MaxAuthTries.*",
			Line:         "MaxAuthTries 3",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^PermitRootLogin.*",
			Line:         "PermitRootLogin prohibit-password",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^PermitEmptyPasswords.*",
			Line:         "PermitEmptyPasswords no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^ChallengeResponseAuthentication.*",
			Line:         "ChallengeResponseAuthentication no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^KerberosAuthentication.*",
			Line:         "KerberosAuthentication no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^GSSAPIAuthentication.*",
			Line:         "GSSAPIAuthentication no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^X11Forwarding.*",
			Line:         "X11Forwarding no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^PermitUserEnvironment.*",
			Line:         "PermitUserEnvironment no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^AllowAgentForwarding.*",
			Line:         "AllowAgentForwarding no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^AllowTcpForwarding.*",
			Line:         "AllowTcpForwarding no",
			State:        "present",
			ShouldCreate: true,
		},
		{
			Regex:        "^PermitTunnel.*",
			Line:         "PermitTunnel no",
			State:        "present",
			ShouldCreate: true,
		},
	}

	sshdFile := File{
		Path:    "/etc/ssh/sshd_config",
		Changes: sshdChanges,
	}

	return HardenConf{
		Files: []File{sshdFile},
	}
}
