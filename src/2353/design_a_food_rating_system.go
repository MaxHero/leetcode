package _2353

import "container/heap"

type Food struct {
	name       string
	rating     int32
	cuisineIdx int16
	cuisinePos int16
}

type Cuisine []*Food

func (c Cuisine) Len() int { return len(c) }
func (c Cuisine) Less(i, j int) bool {
	return c[i].rating > c[j].rating || (c[i].rating == c[j].rating && c[i].name < c[j].name)
}
func (c Cuisine) Swap(i, j int) {
	c[i].cuisinePos = int16(j)
	c[j].cuisinePos = int16(i)
	c[i], c[j] = c[j], c[i]
}
func (c *Cuisine) Push(x interface{}) { *c = append(*c, x.(*Food)) }
func (c *Cuisine) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

type FoodRatings struct {
	foodNameToIdx    map[string]int16
	foods            []Food
	cuisineNameToIdx map[string]int16
	cuisines         []Cuisine
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodsSl := make([]Food, len(foods))
	foodsM := make(map[string]int16, len(foods))
	cuisineM := make(map[string]int16)

	for i := 0; i < len(foods); i++ {
		foodsSl[i].name = foods[i]
		foodsSl[i].rating = int32(ratings[i])
		foodsM[foods[i]] = int16(i)
		cuisineM[cuisines[i]]++
	}

	cuisinesSl := make([]Cuisine, 0, len(cuisineM))
	for cuisine, count := range cuisineM {
		cuisineIdx := int16(len(cuisinesSl))
		cuisinesSl = append(cuisinesSl, make([]*Food, 0, count))
		cuisineM[cuisine] = cuisineIdx
	}

	for i := 0; i < len(foods); i++ {
		cuisineIdx := cuisineM[cuisines[i]]
		foodsSl[i].cuisineIdx = cuisineIdx
		foodsSl[i].cuisinePos = int16(len(cuisinesSl[cuisineIdx]))
		cuisinesSl[cuisineIdx] = append(cuisinesSl[cuisineIdx], &foodsSl[i])
	}

	for i := 0; i < len(cuisinesSl); i++ {
		heap.Init(&cuisinesSl[i])
	}

	return FoodRatings{
		foodNameToIdx:    foodsM,
		foods:            foodsSl,
		cuisineNameToIdx: cuisineM,
		cuisines:         cuisinesSl,
	}
}

func (f FoodRatings) ChangeRating(foodName string, newRating int) {
	foodIdx := f.foodNameToIdx[foodName]
	if f.foods[foodIdx].rating != int32(newRating) {
		f.foods[foodIdx].rating = int32(newRating)
		heap.Fix(&f.cuisines[f.foods[foodIdx].cuisineIdx], int(f.foods[foodIdx].cuisinePos))
	}
}

func (f FoodRatings) HighestRated(cuisine string) string {
	return f.cuisines[f.cuisineNameToIdx[cuisine]][0].name
}
