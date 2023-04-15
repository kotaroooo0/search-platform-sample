package example.kotaroooo0.searchApi.infrastructure

import okhttp3.MediaType
import okhttp3.MediaType.Companion.toMediaType
import okhttp3.OkHttpClient
import okhttp3.Request
import okhttp3.RequestBody
import okhttp3.RequestBody.Companion.toRequestBody
import org.springframework.beans.factory.annotation.Value
import org.springframework.context.annotation.Configuration

// TODO:実装
@Configuration
class SolrQueryClient(private val client: OkHttpClient = OkHttpClient()) {
    @Value("\${solr.host}")
    private val solrHost: String = "http://localhost:8983"

    private val mediaType: MediaType = "application/json; charset=utf-8".toMediaType()


    fun query(collection: String, json: String): SolrQueryResponse {
        val body: RequestBody = json.toRequestBody(mediaType)
        val request: Request = Request.Builder()
            .url(solrHost)
            .post(body)
            .build()
        client.newCall(request).execute().use { response ->
            return SolrQueryResponse()
        }
    }
}