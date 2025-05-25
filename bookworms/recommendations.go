package main

// func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
// 	sb := make(bookRecommendations)

// 	// Register all books on everyone's shelf.
// 	for _, bookworm := range bookworms {
// 		for i, book := range bookworm.Books {
// 			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
// 			registerBookRecommendations(sb, book, otherBooksOnShelves)
// 		}
// 	}

// 	// Recommend a list of related books to each bookworm.
// 	recommendations := make([]Bookworm, len(bookworms))
// 	for i, bookworm := range bookworms {
// 		recommendations[i] = Bookworm{
// 			Name:  bookworm.Name,
// 			Books: recommendBooks(sb, bookworm.Books),
// 		}
// 	}

// 	return recommendations
// }
