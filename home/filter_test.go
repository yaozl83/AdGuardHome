package home

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilters(t *testing.T) {
	dir := prepareTestDir()
	_ = os.MkdirAll(dir+"/data/"+filterDir, 0755)
	defer func() { _ = os.RemoveAll(dir) }()
	Context = homeContext{}
	Context.workDir = dir
	Context.client = &http.Client{
		Timeout: 5 * time.Second,
	}

	f := filter{
		URL: "https://adguardteam.github.io/AdGuardSDNSFilter/Filters/filter.txt",
	}

	// download
	ok, err := f.update()
	assert.Equal(t, nil, err)
	assert.True(t, ok)

	// refresh
	ok, err = f.update()
	assert.True(t, !ok && err == nil)

	err = f.load()
	assert.True(t, err == nil)

	f.unload()
	_ = os.Remove(f.Path())
}
