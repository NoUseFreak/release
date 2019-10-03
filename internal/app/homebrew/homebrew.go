package homebrew

import (
	"fmt"
	"path"

	"github.com/NoUseFreak/letitgo/internal/app/utils"
)

func Execute(c Config) error {
	setDefaults(&c)
	templateProps(&c)
	hash, err := utils.BuildURLHash("sha256", c.URL)
	if err != nil {
		return err
	}
	c.Hash = hash

	content, err := utils.Template(homebrewTpl, c)
	if err != nil {
		return err
	}

	filename := path.Join(c.Folder, fmt.Sprintf("%s.rb", c.Name))
	message := fmt.Sprintf("Upgrade %s to %s", c.Name, c.Version)

	return utils.PublishFile(c.Tap.URL, filename, content, message)
}

func setDefaults(c *Config) {
	if c.Folder == "" {
		c.Folder = "Formula"
	}
	if c.Install == "" {
		c.Install = "bin.install \"{{ .Name }}\""
	}
}

func templateProps(c *Config) {
	utils.TemplateProperty(&c.URL, c)
	utils.TemplateProperty(&c.Install, c)
	utils.TemplateProperty(&c.Test, c)
}
