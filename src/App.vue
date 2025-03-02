<script setup lang="ts">
import ThemeSelection from './components/ThemeSelection.vue';
import SlideshowPage from './pages/SlideshowPage.vue';
import { useNotebookStore, useSettingsStore } from './stores';

const data = useNotebookStore()
data.loadNotebook()
</script>

<template>
    <ThemeSelection v-if="useSettingsStore().isSelectingTheme" />
    <SlideshowPage v-if="data.notebookExists" />
</template>

<style scoped></style>
