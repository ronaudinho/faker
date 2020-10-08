package faker

// UserAgent is a faker struct for UserAgent
type UserAgent struct {
	Faker *Faker
}

// InternetExplorer returns a fake InternetExplorer UserAgent
func (u UserAgent) InternetExplorer() string {
	return "Mozilla/5.0 (compatible; MSIE 7.0; Windows 98; Win 9x 4.90; Trident/3.0)"
}

// Opera returns a fake Opera UserAgent
func (u UserAgent) Opera() string {
	return "Opera/8.25 (Windows NT 5.1; en-US) Presto/2.9.188 Version/10.00"
}

// Safari returns a fake Safari UserAgent
func (u UserAgent) Safari() string {
	return "Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_7_1 rv:3.0; en-US) AppleWebKit/534.11.3 (KHTML, like Gecko) Version/4.0 Safari/534.11.3"
}

// Firefox returns a fake Firefox UserAgent
func (u UserAgent) Firefox() string {
	return "Mozilla/5.0 (X11; Linuxi686; rv:7.0) Gecko/20101231 Firefox/3.6"
}

// Chrome returns a fake Chrome UserAgent
func (u UserAgent) Chrome() string {
	return "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) AppleWebKit/5312 (KHTML, like Gecko) Chrome/14.0.894.0 Safari/5312"
}

// UserAgent returns a fake UserAgent
func (u UserAgent) UserAgent() string {
	agents := []string{
		u.InternetExplorer(),
		u.Opera(),
		u.Safari(),
		u.Firefox(),
		u.Chrome(),
	}

	return u.Faker.RandomStringElement(agents)
}
