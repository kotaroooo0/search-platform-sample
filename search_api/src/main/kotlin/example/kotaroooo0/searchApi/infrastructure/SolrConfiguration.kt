package example.kotaroooo0.searchApi.infrastructure

import org.apache.solr.client.solrj.SolrClient
import org.apache.solr.client.solrj.impl.Http2SolrClient
import org.springframework.beans.factory.annotation.Value
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration


@Configuration
class SolrConfiguration {

    @Value("\${solr.host}")
    val solrHost: String = "http://localhost:8983"

    @Bean
    fun solrClient(): SolrClient {
        return Http2SolrClient.Builder(solrHost).build()
    }
}