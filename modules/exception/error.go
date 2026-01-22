package exception

func CreateError(err error) map[string]string {
	return map[string]string{
		"error": err.Error(),
	}
}
