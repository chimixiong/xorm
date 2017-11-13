package xorm

import (
	"html/template"
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet"
	"gopkg.in/flosch/pongo2.v3"
)

type SqlTemplate interface {
	WalkFunc(path string, info os.FileInfo, err error) error
	paresSqlTemplate(filename string, filepath string) error
	ReadTemplate(filepath string) ([]byte, error)
	Execute(key string, args ...interface{}) (string, error)
	RootDir() string
	Extension() string
	SetSqlTemplateCipher(cipher Cipher)
	LoadSqlTemplate(filepath string) error
	BatchLoadSqlTemplate(filepathSlice []string) error
	ReloadSqlTemplate(filepath string) error
	BatchReloadSqlTemplate(filepathSlice []string) error
	AddSqlTemplate(key string, sqlTemplateStr string) error
	UpdateSqlTemplate(key string, sqlTemplateStr string) error
	RemoveSqlTemplate(key string)
	BatchAddSqlTemplate(key string, sqlTemplateStrMap map[string]string) error
	BatchUpdateSqlTemplate(key string, sqlTemplateStrMap map[string]string) error
	BatchRemoveSqlTemplate(key []string)
}

func Pongo2(directory, extension string) *Pongo2Template {
	template := make(map[string]*pongo2.Template, 100)
	return &Pongo2Template{
		SqlTemplateRootDir: directory,
		extension:          extension,
		Template:           template,
	}
}

func Default(directory, extension string) *HTMLTemplate {
	template := make(map[string]*template.Template, 100)
	return &HTMLTemplate{
		SqlTemplateRootDir: directory,
		extension:          extension,
		Template:           template,
	}
}

func Jet(directory, extension string) *JetTemplate {
	template := make(map[string]*jet.Template, 100)
	return &JetTemplate{
		SqlTemplateRootDir: directory,
		extension:          extension,
		Template:           template,
	}
}

func (engine *Engine) RegisterSqlTemplate(sqlt SqlTemplate, Cipher ...Cipher) error {
	engine.SqlTemplate = sqlt
	if len(Cipher) > 0 {
		engine.SqlTemplate.SetSqlTemplateCipher(Cipher[0])
	}
	err := filepath.Walk(engine.SqlTemplate.RootDir(), engine.SqlTemplate.WalkFunc)
	if err != nil {
		return err
	}

	return nil
}

func (engine *Engine) LoadSqlTemplate(filepath string) error {
	return engine.SqlTemplate.LoadSqlTemplate(filepath)
}

func (engine *Engine) BatchLoadSqlTemplate(filepathSlice []string) error {
	return engine.SqlTemplate.BatchLoadSqlTemplate(filepathSlice)
}

func (engine *Engine) ReloadSqlTemplate(filepath string) error {
	return engine.SqlTemplate.ReloadSqlTemplate(filepath)
}

func (engine *Engine) BatchReloadSqlTemplate(filepathSlice []string) error {
	return engine.SqlTemplate.BatchReloadSqlTemplate(filepathSlice)
}

func (engine *Engine) AddSqlTemplate(key string, sqlTemplateStr string) error {
	return engine.SqlTemplate.AddSqlTemplate(key, sqlTemplateStr)
}

func (engine *Engine) UpdateSqlTemplate(key string, sqlTemplateStr string) error {
	return engine.SqlTemplate.UpdateSqlTemplate(key, sqlTemplateStr)
}

func (engine *Engine) RemoveSqlTemplate(key string) {
	engine.SqlTemplate.RemoveSqlTemplate(key)
}

func (engine *Engine) BatchAddSqlTemplate(key string, sqlTemplateStrMap map[string]string) error {
	return engine.SqlTemplate.BatchAddSqlTemplate(key, sqlTemplateStrMap)

}

func (engine *Engine) BatchUpdateSqlTemplate(key string, sqlTemplateStrMap map[string]string) error {
	return engine.SqlTemplate.BatchUpdateSqlTemplate(key, sqlTemplateStrMap)

}

func (engine *Engine) BatchRemoveSqlTemplate(key []string) {
	engine.SqlTemplate.BatchRemoveSqlTemplate(key)
}
