package example.kotaroooo0.searchApi.infrastructure

import example.kotaroooo0.searchApi.domain.objects.books.BookId
import org.apache.solr.client.solrj.beans.Field

data class BookDocument(
    @Field("book_id")
    val bookId: BookId
)
