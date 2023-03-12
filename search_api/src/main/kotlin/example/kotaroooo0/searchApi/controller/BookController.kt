package example.kotaroooo0.searchApi.controller

import example.kotaroooo0.searchApi.domain.service.BookService
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/book")
class BookController (private val bookService: BookService) {

    @GetMapping("/search")
    fun getSearch(@RequestParam("keyword", required = true) keyword: String): BookSearchResponse {
        val bookIdList = bookService.getSearch(keyword).map {
            it.id
        }
        return BookSearchResponse(bookIdList)
    }
}