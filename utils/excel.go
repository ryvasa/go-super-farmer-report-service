package utils

import "sync"

var (
	// Map untuk menyimpan path file Excel berdasarkan ID request
	filePathMap = make(map[string]string)
	mutex       sync.RWMutex
)

// SetFilePath menyimpan path file untuk ID tertentu
func SetFilePath(requestID string, path string) {
	mutex.Lock()
	defer mutex.Unlock()
	filePathMap[requestID] = path
}

// GetFilePath mengambil path file berdasarkan ID
func GetFilePath(requestID string) (string, bool) {
	mutex.RLock()
	defer mutex.RUnlock()
	path, exists := filePathMap[requestID]
	return path, exists
}

// RemoveFilePath menghapus entry setelah file didownload
func RemoveFilePath(requestID string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(filePathMap, requestID)
}
