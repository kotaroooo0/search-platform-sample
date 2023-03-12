package example.kotaroooo0.searchApi.domain.objects.books

data class SearchParameter(
    val keyword: String,
    val isMens: Boolean, // メンズ向けかどうか
    val isNew: Boolean, // 新作かどうか
)
