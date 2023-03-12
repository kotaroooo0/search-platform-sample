package example.kotaroooo0.searchApi.domain.repository

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.objects.books.SearchParameter

interface BookRepository {
    fun search(searchParameter: SearchParameter): List<BookId>
}

