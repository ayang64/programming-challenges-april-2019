package domains

// given a slice of strings, return a comma separated string containing all of
// the unique domains assocated with either HTTP or HTTPS URLS.
//
// the domains in the result list should be squashed to lower case
//
// if the domain starts with www., w3., or web., that part of it should be
// stripped from the domain.
//
// For example, if the slice contains:
//
//  hi, i typically like to search <a href='https://www.google.com/'>http://google.com/!</a>
//  for information but sometimes i enjoy using http://www.duckduckgo.com/
//
//  someitmes it returns results from http://golang.org and https://reddit.com/
//
// the resulting slice should contain:
//
//  duckduckgo.com,golang.org,google.com,reddit.com"
//
func ExtractDomains(s []string) []string {
	// insert your code here.
	return nil
}
