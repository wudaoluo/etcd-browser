package etcdlib



import (
	"path"
	"strings"
)

func isRoot(key string) bool {
	return key == "/"
}


// ensure key, return (realKey, parentKey)
func (c *client) ensureKey(key string) (string, string, error) {
	if !strings.HasPrefix(key, "/") {
		return "", "", ErrorInvalidKey
	}

	if isRoot(key) {
		return c.prefix, c.prefix, nil
	} else {
		realKey := c.prefix + key
		return realKey, path.Clean(realKey + "/../"), nil
	}
}
