package utils

import (
	"fmt"
	"strings"
)

func QueueNameUtils(queueNamePrefix, queueNameSuffix, key string) string {
	return fmt.Sprintf("%s%s%s", queueNamePrefix, strings.ReplaceAll(key, " ", ""), queueNameSuffix)
}

func QueueList(queueList string) []string {
	return strings.Split(queueList, ",")
}
