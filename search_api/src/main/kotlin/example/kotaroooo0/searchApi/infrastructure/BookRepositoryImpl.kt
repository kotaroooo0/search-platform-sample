package example.kotaroooo0.searchApi.infrastructure

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.objects.books.SearchParameter
import example.kotaroooo0.searchApi.domain.repository.BookRepository
import org.springframework.stereotype.Repository

@Repository
class BookRepositoryImpl : BookRepository {
    override fun search(searchParameter: SearchParameter): List<BookId> {
        // TODO("Not yet implemented")
        return listOf(
            BookId("1"),
            BookId("2"),
            BookId("3"),
            BookId("4"),
        )
    }
}