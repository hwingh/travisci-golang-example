package fizzbuzz

import (
	"strconv"
)

type FizzBuzzHandler struct {
	cache Cache
}

func (h *FizzBuzzHandler) RunFizzBuzz(args []string) ([]string, error) {

	nums, err := converInput(args)
	if err != nil {
		return nil, err
	}

	var rsp []string

	for _, num := range nums {
		str, ok := h.cache.Get(num)

		if !ok {
			str = fizzBuzz(num)
			h.cache.Put(num, str)
		}

		rsp = append(rsp, str)
	}

	return rsp, nil
}

func NewHandler(c Cache) *FizzBuzzHandler {
	return &FizzBuzzHandler{
		cache: c,
	}
}

func converInput(args []string) ([]int, error) {
	var rsp []int

	for _, str := range args {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		rsp = append(rsp, num)
	}

	return rsp, nil
}
