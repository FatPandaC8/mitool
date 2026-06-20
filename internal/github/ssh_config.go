package github

import (
	"fmt"
	"os"
	"strings"

	"github.com/FatPandaC8/mitool/internal/config"
)

func AddSSHConfig(
	name string,
	keyPath string,
) error {

	host := "github-" + name

	block := fmt.Sprintf(
		`# BEGIN MITOOL %s
		Host %s
			HostName github.com
			User git
			IdentityFile %s
		# END MITOOL %s
		`,
		name,
		host,
		keyPath,
		name,
	)


	path := config.SSHConfigFile()


	data, err := os.ReadFile(path)

	if err != nil && !os.IsNotExist(err) {
		return err
	}


	content := string(data)


	// avoid duplicate blocks
	if strings.Contains(
		content,
		"BEGIN MITOOL "+name,
	) {
		return nil
	}


	if len(content) > 0 &&
		!strings.HasSuffix(content, "\n") {

		content += "\n"
	}


	content += block


	return os.WriteFile(
		path,
		[]byte(content),
		0600,
	)
}

func RemoveSSHConfig(name string) error {
	path := config.SSHConfigFile()

	data, err := os.ReadFile(path)

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}


	start := "# BEGIN MITOOL " + name
	end := "# END MITOOL " + name


	content := string(data)

	startIndex := strings.Index(content, start)

	if startIndex == -1 {
		return nil
	}


	endIndex := strings.Index(
		content[startIndex:],
		end,
	)

	if endIndex == -1 {
		return nil
	}


	endIndex += startIndex + len(end)


	newContent := content[:startIndex] +
		content[endIndex:]


	return os.WriteFile(
		path,
		[]byte(newContent),
		0600,
	)
}