// Package anagram serves to check anagrams
package anagram

import (
	"sort"
	"strings"
)

// Detect check anagrams
func DetectV1(subject string, candidates []string) []string {
	var ret []string
	inp := make(map[rune]int, len(subject))
	sub := make(map[rune]int, len(subject))

	lowSubject := strings.ToLower(subject)
	for _, r := range strings.ToLower(subject) {
		inp[r]++
	}

	for _, candidat := range candidates {
		lowCandidat := strings.ToLower(candidat)
		if (len(subject) != len(candidat)) || (lowSubject == lowCandidat) {
			continue
		}

		for r := range inp {
			sub[r] = inp[r]
		}

		ok := true
		for _, r := range lowCandidat {
			if sub[r] == 0 {
				ok = false
				break
			}
			sub[r]--
		}
		if ok {
			ret = append(ret, candidat)
		}
	}
	return ret
}
func Detect(subject string, candidates []string) []string {
	var ret []string

	lowSubj := strings.ToLower(subject)
	subj := strings.Split(lowSubj, "")
	sort.Strings(subj)
	subject = strings.Join(subj, "")

	for _, candidate := range candidates {
		lowCand := strings.ToLower(candidate)
		if lowSubj == lowCand {
			return nil
		}
		cand := strings.Split(lowCand, "")
		sort.Strings(cand)
		str := strings.Join(cand, "")
		if subject == str {
			ret = append(ret, candidate)
		}
	}

	return ret
}

func DetectV3(subject string, candidates []string) []string {
	var ret []string

	lowSubj := strings.ToLower(subject)
	subj := strings.Split(lowSubj, "")
	sort.Strings(subj)

	for _, candidate := range candidates {
		lowCand := strings.ToLower(candidate)
		if lowSubj == lowCand || len(lowSubj) != len(lowCand) {
			return nil
		}
		cand := strings.Split(lowCand, "")
		sort.Strings(cand)
		for i, _ := range subj {
			if subj[i] != cand[i] {
				break
			}
		}
		ret = append(ret, candidate)
	}

	return ret
}
