<script setup lang="ts">
import type { ComputedRef } from 'vue';

type Item = {
    label: string | ComputedRef<string> | 'separator',
    action?: () => void
}
const { items } = defineProps<{ items: Item[] }>()
</script>
<template>
    <div class="ContextMenu">
        <template v-for="item in items">
            <button class="MenuItem" v-if="item.label !== 'separator'"
                @click="item.action">
                {{ item.label }}
            </button>
            <div class="sep" v-else />
        </template>
    </div>
</template>
<style scoped lang="scss">
.ContextMenu {
    display: none;
    border-radius: 15px;
    background: rgb(var(--bg-1));
    flex-direction: column;
    gap: 2px;
    padding: 8px;
    position: absolute;
    bottom: 100%;
    left: 50%;
    transform: translateX(-50%);
    opacity: 0;
    transition: opacity .25s;
    border: 1px solid rgb(var(--bg-2));
    box-shadow: var(--shadow-lg);
    text-align: left;
    min-width: 200px;
    cursor: default;
    z-index: 10;

    .sep {
        border: 1px solid rgb(var(--bg-2));
        width: 100%;
        height: 0;
        margin-block: 6px;
    }
}

.MenuItem {
    background: none;
    padding: 8px 12px;
    transition: .15s;
    border-radius: 8px;
    text-align: left;
    font-weight: normal;

    &:hover {
        background: rgb(var(--bg-2));
    }
}
</style>
<style lang="scss">
.FlyoutOpener {
    position: relative;

    &:focus-within .ContextMenu {
        display: flex;
        opacity: 1;
    }
}
</style>