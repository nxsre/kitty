package common

type Pager struct {
	Data       interface{}
	PageNo     int
	PageSize   int
	TotalCount int
	TotalPage  int
}

func (this *Pager)SetTotalCount(count int) {

	this.TotalCount = count
	if this.TotalCount % this.PageSize == 0 {
		this.TotalPage = this.TotalCount / this.PageSize
	} else {
		this.TotalPage = (this.TotalCount / this.PageSize + 1)
	}
}