package models

import (
	"errors"
	"net/url"
)

const errorOnInvalidChar = "custom key contains invalid characters"

var KeyRunes = []rune{'0','1','2','3','4','5','6','7','8','9',
	'A','B','C','D','E','F','G','H','I','J','K','L','M',
	'N','O','P','Q','R','S','T','U','V','W','X','Y','Z',
	'a','b','c','d','e','f','g','h','i','j','k','l','m',
	'n','o','p','q','r','s','t','u','v','w','x','y','z',
}

var mapKeyRunes = map[rune]interface{}{'0':nil,'1':nil,'2':nil,
	'3':nil,'4':nil,'5':nil,'6':nil,'7':nil,'8':nil,'9':nil,
	'A':nil,'B':nil,'C':nil,'D':nil,'E':nil,'F':nil,'G':nil,
	'H':nil,'I':nil,'J':nil,'K':nil,'L':nil,'M':nil,'N':nil,
	'O':nil,'P':nil,'Q':nil,'R':nil,'S':nil,'T':nil,'U':nil,
	'V':nil,'W':nil,'X':nil,'Y':nil,'Z':nil,
	'a':nil,'b':nil,'c':nil,'d':nil,'e':nil,'f':nil,'g':nil,
	'h':nil,'i':nil,'j':nil,'k':nil,'l':nil,'m':nil,'n':nil,
	'o':nil,'p':nil,'q':nil,'r':nil,'s':nil,'t':nil,'u':nil,
	'v':nil,'w':nil,'x':nil,'y':nil,'z':nil,
}

func (s *SaveURLReq) Validate() error{
	_, err := url.Parse(s.OriginalURL)
	if err != nil{
		return err
	}
	return checkKey(s.CustomKey)
}

func (r *RedirectReq) Validate() error{
	return checkKey(r.Key)
}

func checkKey(key string) error{
	for _, r := range key{
		if _, ok := mapKeyRunes[r]; !ok{
			return errors.New(errorOnInvalidChar)
		}
	}
	return nil
}