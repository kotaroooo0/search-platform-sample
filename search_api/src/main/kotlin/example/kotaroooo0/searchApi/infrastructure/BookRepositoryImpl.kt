package example.kotaroooo0.searchApi.infrastructure

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.objects.books.SearchParameter
import example.kotaroooo0.searchApi.domain.repository.BookRepository
import org.springframework.stereotype.Repository

@Repository
class BookRepositoryImpl(private val solrClient: SolrQueryClient) : BookRepository {

    companion object {
        private const val BOOK_COLLECTION_NAME = "books"
    }

    override fun search(searchParameter: SearchParameter): List<BookId> {
        // TODO: 実装
        val response = solrClient.query(BOOK_COLLECTION_NAME, "")
        return listOf()
    }
}
