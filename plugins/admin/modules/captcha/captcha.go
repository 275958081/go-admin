package captcha

type Captcha interface {
	Validate(token string) bool
}

var List = make(map[string]Captcha)

func Add(key string, captcha Captcha) {
	List[key] = captcha
}

func Get(key string) (Captcha, bool) {
	captcha, ok := List[key]
	return captcha, ok
}
