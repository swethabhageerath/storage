package redis

type ErrSetRedisCache struct {
	Code    int
	Message string
}

func NewErrSetRedisCache() ErrSetRedisCache {
	return ErrSetRedisCache{
		Code:    500,
		Message: "Error setting value in Redis Cache",
	}
}

func (e ErrSetRedisCache) Error() string {
	return e.Message
}

type ErrKeyNotExists struct {
	Code    int
	Message string
}

func NewErrKeyNotExists() ErrKeyNotExists {
	return ErrKeyNotExists{
		Code:    500,
		Message: "Specified key %s does not exist in redis cache",
	}
}

func (e ErrKeyNotExists) Error() string {
	return e.Message
}

type ErrUnknownError struct {
	Code    int
	Message string
}

func NewErrUnknownError() ErrUnknownError {
	return ErrUnknownError{
		Code:    500,
		Message: "Unknown error occuring while performing your request",
	}
}

func (e ErrUnknownError) Error() string {
	return e.Message
}

type ErrRetrievingKey struct {
	Code    int
	Message string
}

func NewErrRetrievingKey() ErrRetrievingKey {
	return ErrRetrievingKey{
		Code:    500,
		Message: "Error retrieving key %s from redis cache",
	}
}

func (e ErrRetrievingKey) Error() string {
	return e.Message
}
