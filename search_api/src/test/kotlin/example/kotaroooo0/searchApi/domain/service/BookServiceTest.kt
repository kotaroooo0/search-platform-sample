package example.kotaroooo0.searchApi.domain.service

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.objects.books.SearchParameter
import example.kotaroooo0.searchApi.domain.repository.BookRepository
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import org.mockito.Mockito.mock
import org.mockito.Mockito.`when`

class BookServiceTest {
    private val bookRepository = mock(BookRepository::class.java)
    private val bookService = BookService(bookRepository)

    @Test
    fun testGetSearch() {
        val keyword = "kotlin"
        val bookIds = listOf(BookId("12345"), BookId("67890"))
        `when`(bookRepository.search(SearchParameter(keyword, isNew = true, isMens = true)))
            .thenReturn(bookIds)

        val result = bookService.getSearch(keyword)

        assertEquals(bookIds, result)
    }
}
