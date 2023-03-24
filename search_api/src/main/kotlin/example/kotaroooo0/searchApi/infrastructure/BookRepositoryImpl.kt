package example.kotaroooo0.searchApi.infrastructure

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import example.kotaroooo0.searchApi.domain.objects.books.SearchParameter
import example.kotaroooo0.searchApi.domain.repository.BookRepository
import org.apache.solr.client.solrj.SolrClient
import org.apache.solr.client.solrj.beans.DocumentObjectBinder
import org.apache.solr.client.solrj.response.QueryResponse
import org.apache.solr.common.params.MapSolrParams
import org.springframework.stereotype.Repository

@Repository
class BookRepositoryImpl(private val solrClient: SolrClient) : BookRepository {

    companion object {
        private const val BOOK_COLLECTION_NAME = "books"
    }

    override fun search(searchParameter: SearchParameter): List<BookId> {
        // TODO: searchParameterを使うクエリを書く
        val queryParams = MapSolrParams(mapOf(
            "q" to "*:*",
            "rows" to "5"
        ))

        val response: QueryResponse = solrClient.query(BOOK_COLLECTION_NAME, queryParams)
        val bookDocuments: List<BookDocument> = DocumentObjectBinder().getBeans(BookDocument::class.java, response.results)
        return bookDocuments.map { it.bookId }
    }
}
