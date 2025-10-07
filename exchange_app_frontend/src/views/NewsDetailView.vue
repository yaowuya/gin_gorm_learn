<script setup lang="ts">
import { onMounted, ref } from 'vue'
import type { Article } from '@/types/Article'
import { useRoute } from 'vue-router'
import axios from '@/axios.ts'

const article = ref<Article | null>(null)
const route = useRoute()

const fetchArticle = async () => {
  const { id } = route.params
  try {
    const response = await axios.get<Article>(`/articles/${id}`)
    article.value = response.data
  } catch (err) {
    console.log('Failed to load article', err)
  }
}

onMounted(fetchArticle)
</script>

<template>
  <el-container>
    <el-main>
      <el-card v-if="article" class="article-detail">
        <h1>{{ article.Title }}</h1>
        <p>{{ article.Content }}</p>
      </el-card>
      <div v-else class="no-data">文章不存在</div>
    </el-main>
  </el-container>
</template>

<style scoped>
.article-detail {
  margin: 20px 0;
}

.no-data {
  text-align: center;
  font-size: 1.2em;
  color: #999;
}
</style>
