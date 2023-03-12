package example.kotaroooo0.searchApi.domain.service

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.objects.books.SearchParameter
import example.kotaroooo0.searchApi.domain.repository.BookRepository
import org.springframework.stereotype.Service

@Service
class BookService(private val bookRepository: BookRepository) {
    fun getSearch(keyword: String): List<BookId>{
        return bookRepository.search(rewriteSearchQuery(keyword))
    }

    private fun rewriteSearchQuery(keyword: String): SearchParameter{
        return SearchParameter(keyword, isNew = true, isMens = true)
    }
}