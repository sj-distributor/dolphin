/*
 * @Author: Marlon.M
 * @Email: maiguangyang@163.com
 * @Date: 2024-09-24 17:15:17
 */
package utils

import "os"

func CreateDirIfNotExists(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0777)
		if err != nil {
			panic(err)
		}
	}
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		return true
	}
	return false
}
