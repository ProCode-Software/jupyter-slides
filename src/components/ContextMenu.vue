<script setup lang="ts">
import type { ComputedRef } from 'vue'

type Item = {
    label: string | ComputedRef<string> | 'separator'
    action?: () => void
}
const { items } = defineProps<{ items: Item[] }>()
</script>
<template>
    <div class="ContextMenu">
        <template v-for="item in items">
            <button
                class="MenuItem"
                v-if="item.label !== 'separator'"
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
    padding: 6px;
    position: absolute;
    bottom: calc(100% + 10px);
    left: 50%;
    transform: translateX(-50%);
    animation: open 0.3s cubic-bezier(0.165, 0.84, 0.44, 1);
    border: 1px solid rgb(var(--bg-2));
    box-shadow: var(--shadow-lg);
    text-align: left;
    min-width: 220px;
    cursor: default;
    z-index: 10;
    max-height: 70vh;
    overflow-y: auto;

    .sep {
        border: 1px solid rgb(var(--bg-2));
        width: 100%;
        height: 0;
        margin-block: 4px;
    }
}

@keyframes open {
    from {
        transform: translateX(-50%) translateY(15px);
    }

    to {
        transform: translateX(-50%) translateY(0);
    }
}

.MenuItem {
    background: none;
    padding: 8px 12px;
    transition: background 0.15s;
    border-radius: 8px;
    text-align: left;
    font-weight: normal;

    &:hover {
        background: rgb(var(--bg-2));
    }
}

@media (prefers-color-scheme: light) {
    .ContextMenu {
        background: rgb(var(--bg-0));
    }

    .MenuItem:hover {
        background: rgb(var(--bg-1));
    }
}
</style>
<style lang="scss">
.FlyoutOpener {
    position: relative;

    &:focus-within .ContextMenu {
        display: flex;
    }
}
</style>
