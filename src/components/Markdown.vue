<script setup lang="ts">
import { useMarkdown } from '@/composables/useMarkdown';

const { content } = defineProps<{ content: string }>()
</script>
<template>
    <div class="Markdown" v-html="useMarkdown(content)" />
</template>