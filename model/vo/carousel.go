package vo

import "github.com/airdb/bbs-api/model/po"

type ListCarouselReq struct {
}

type ListCarouselResp struct {
	Square []Square
}

type Carousel struct {
	ID     uint   `json:"id"`
	Href   string `json:"href"`
	Image  string `json:"image"`
	Title  string `json:"title"`
	Flag   uint   `json:"flag"`
	Status uint   `json:"status"`
}

func FromPoCarousel(pCarousel *po.Carousel) *Carousel {
	return &Carousel{
		ID:     pCarousel.ID,
		Href:   pCarousel.Href,
		Title:  pCarousel.Title,
		Image:  pCarousel.Image,
		Flag:   pCarousel.Flag,
		Status: pCarousel.Status,
	}
}

func FromPoCarousels(pCarousels []*po.Carousel) (vCarousels []*Carousel) {
	for _, notice := range pCarousels {
		vCarousels = append(vCarousels, FromPoCarousel(notice))
	}

	return
}

func ListCarousel() []*Carousel {
	return FromPoCarousels(po.ListCarousel())
}
