package timestamp

// (Tabi[i].date, i) <2 (Tabi[k].date, k)
func compare_timestamp(date_1, num_site_1, date_2, num_site_2 int) bool {
	return date_1 < date_2 || (date_1 == date_2 && num_site_1 < num_site_2)
}
