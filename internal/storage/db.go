package storage

import (
	"archive/zip"
	"io/ioutil"
	"os"
	"path"
	"sync"
)

// Storage structure
type Storage struct {
	s   Settings
	rwm sync.RWMutex
}

// Settings structure
type Settings struct {
	Path string
}

// New creates a storage object
func New(s Settings) *Storage {

	err := os.MkdirAll(s.Path, 0600)
	if err != nil {
		panic(err)
	}

	return &Storage{
		s: s,
	}
}

// Get function
func (s *Storage) Get(k string) ([]byte, error) {
	s.rwm.RLock()
	defer s.rwm.RUnlock()

	filePath := path.Join(s.s.Path, k)
	b, err := ioutil.ReadFile(filePath)

	return b, err
}

// Put function
func (s *Storage) Put(k string, v []byte) error {

	s.rwm.Lock()
	defer s.rwm.Unlock()

	filePath := path.Join(s.s.Path, k)
	err := ioutil.WriteFile(filePath, v, 0600)

	return err
}

// Delete function
func (s *Storage) Delete(k string) error {
	s.rwm.Lock()
	defer s.rwm.Unlock()

	filePath := path.Join(s.s.Path, k)
	return os.Remove(filePath)

}

// GetIndex ...
func (s *Storage) GetIndex() ([]os.FileInfo, error) {

	s.rwm.Lock()
	defer s.rwm.Unlock()

	files, err := ioutil.ReadDir(s.s.Path)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// BackupTo ...
func (s *Storage) BackupTo(dest string) error {

	// GetIndex is atomic
	files, err := s.GetIndex()
	if err != nil {
		return err
	}

	// request lock here
	s.rwm.Lock()
	defer s.rwm.Unlock()

	// Create a buffer to write our archive to.
	newfile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer newfile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(newfile)

	// Add some files to the archive.
	for _, file := range files {
		fn := file.Name()

		f, err := w.Create(fn)
		if err != nil {
			return err
		}

		rb, _ := ioutil.ReadFile(path.Join(s.s.Path, fn))
		_, err = f.Write(rb)
		if err != nil {
			return err
		}
	}

	return w.Close()
}
