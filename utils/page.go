package utils

import "math"

func Rangelist(a, b int, c []int) []int {
	for i := a; i < b; i++ {
		c = append(c, i)
	}
	return c

}

func GetPageNum(count int64, pagesize int) int {
	pagenum := int(math.Ceil(float64(count) / float64(pagesize)))
	return pagenum
}

// ... 6 7 8 9 当前页 11 12 13 14 ...

//总页数,当前页，前后页数
func Get_pagination_data(num_pages, current_page, around_count int) (left_pages, right_pages []int, left_has_more, right_has_more bool) {

	if current_page <= around_count+2 {
		left_pages = Rangelist(1, current_page, left_pages)
	} else {
		left_has_more = true
		left_pages = Rangelist(current_page-around_count, current_page, left_pages)
	}

	if current_page >= num_pages-around_count-1 {
		right_pages = Rangelist(current_page+1, num_pages+1, right_pages)
	} else {
		right_has_more = true
		right_pages = Rangelist(current_page+1, current_page+around_count+1, right_pages)

	}
	return
}

func HasNext(page, num_pages int) (has_previous, has_next bool, previous_page_number, next_page_number int) {
	if page == 1 && num_pages == page {
		has_previous = false
		has_next = false
		next_page_number = 1
		previous_page_number = 1
	} else if page == 1 && num_pages > page {
		has_previous = false
		has_next = true
		previous_page_number = page
		next_page_number = page + 1
	} else if page > 1 && num_pages > page {
		has_previous = true
		has_next = true
		previous_page_number = page - 1
		next_page_number = page + 1
	} else if page > 1 && page == num_pages {
		has_previous = true
		has_next = false
		previous_page_number = page - 1
		next_page_number = page
	}
	return
}
