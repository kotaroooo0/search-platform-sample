package example.kotaroooo0.searchApi.controller

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.service.BookService
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test
import org.junit.jupiter.api.extension.ExtendWith
import org.mockito.ArgumentMatchers.anyString
import org.mockito.Mock
import org.mockito.Mockito.`when`
import org.mockito.junit.jupiter.MockitoExtension

@ExtendWith(MockitoExtension::class)
class BookControllerTest {

    @Mock
    private lateinit var bookService: BookService

    @Test
    fun `getSearch returns BookSearchResponse with bookIdList`() {
        val bookIdList = listOf(BookId("id1"), BookId("id2"), BookId("id3"))
        `when`(bookService.getSearch(anyString())).thenReturn(bookIdList)

        val controller = BookController(bookService)

        val actual = controller.getSearch("keyword")

        assertEquals(bookIdList.map { it.id }, actual.bookIds)
    }
}
