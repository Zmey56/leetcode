package numUniqueEmails

//Every valid email consists of a local name and a domain name, separated by the '@' sign. Besides lowercase letters,
//the email may contain one or more '.' or '+'.
//
//For example, in "alice@leetcode.com", "alice" is the local name, and "leetcode.com" is the domain name.
//If you add periods '.' between some characters in the local name part of an email address, mail sent there will be
//forwarded to the same address without dots in the local name. Note that this rule does not apply to domain names.
//
//For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.
//If you add a plus '+' in the local name, everything after the first plus sign will be ignored. This allows certain
//emails to be filtered. Note that this rule does not apply to domain names.
//
//For example, "m.y+name@email.com" will be forwarded to "my@email.com".
//It is possible to use both of these rules at the same time.
//
//Given an array of strings emails where we send one email to each emails[i], return the number of different addresses that actually receive mails.

func numUniqueEmails(emails []string) int {
	uniqueEmails := make(map[string]struct{})
	for _, email := range emails {
		localName, domainName := splitEmail(email)
		uniqueEmails[localName+domainName] = struct{}{}
	}
	return len(uniqueEmails)
}

func splitEmail(email string) (string, string) {
	atIndex := 0
	for i, v := range email {
		if v == '@' {
			atIndex = i
			break
		}
	}
	localName := email[:atIndex]
	domainName := email[atIndex:]
	localName = removeDot(localName)
	localName = removePlus(localName)
	return localName, domainName
}

func removeDot(localName string) string {
	var res string
	for _, v := range localName {
		if v != '.' {
			res += string(v)
		}
	}
	return res
}

func removePlus(localName string) string {
	var res string
	for _, v := range localName {
		if v == '+' {
			break
		}
		res += string(v)
	}
	return res
}
