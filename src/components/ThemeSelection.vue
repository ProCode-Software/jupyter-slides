<script setup lang="ts">
import { friendlyTheme } from '@/composables/utils'
import { useSettingsStore } from '@/stores'
import { onKeyStroke } from '@vueuse/core'
import { bundledThemes } from 'shiki'
import { ref } from 'vue'

const themes = Object.keys(bundledThemes)

const { setTheme, hideThemeSelection, theme } = useSettingsStore()
const selection = ref(theme)
const selectTheme = () => setTheme(selection.value)

onKeyStroke('Escape', hideThemeSelection)
</script>
<template>
    <div class="Dialog__Overlay" />
    <div class="Dialog">
        <div class="DialogTitle">Theme</div>
        <div class="DialogContent">
            <div class="Selector">
                <button
                    v-for="theme in themes"
                    @click="selection = theme"
                    key="theme"
                    :class="['item', { selected: selection === theme }]">
                    {{ friendlyTheme(theme) }}
                </button>
            </div>
        </div>
        <div class="DialogButtons">
            <button @click="hideThemeSelection">Cancel</button>
            <button class="primary" @click="selectTheme">OK</button>
        </div>
    </div>
</template>
<style scoped lang="scss">
.Dialog__Overlay {
    position: absolute;
    width: 100%;
    height: 100%;
    background: rgb(0, 0, 0, 0.2);
    left: 0;
    top: 0;
    z-index: 20;
}

.Dialog {
    display: flex;
    border-radius: 20px;
    flex-direction: column;
    background: rgba(var(--bg-0), 0.9);
    backdrop-filter: blur(20px) brightness(1.5);
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border: 1px solid rgb(var(--bg-2));
    box-shadow: var(--shadow-lg);
    width: 50vw;
    z-index: 99;
    animation: open 0.2s cubic-bezier(0.165, 0.84, 0.44, 1);
}

@keyframes open {
    from {
        opacity: 0;
        transform: translate(-50%, -50%) scale(0.7);
    }
    to {
        opacity: 1;
        transform: translate(-50%, -50%) scale(1);
    }
}

.DialogTitle {
    font-size: 1.2em;
    font-weight: 600;
    padding: 15px;
    width: 100%;
    border-bottom: 1px solid rgb(var(--bg-2));
}

.Selector {
    display: flex;
    flex-direction: column;
    gap: 2px;
    width: 100%;
    overflow-y: auto;
    max-height: 50vh;
    padding: 10px;

    .item {
        width: 100%;
        background: none;
        transition: background 0.15s;
        text-align: left;
        font-weight: normal;

        &:hover {
            background: rgb(var(--bg-1));
        }
    }

    .item.selected {
        background: rgb(var(--accent));
        color: #fff;
        font-weight: 500;

        &:hover {
            background: rgb(var(--accent-dark));
        }
    }
}

.DialogButtons {
    width: 100%;
    display: flex;
    justify-content: flex-end;
    padding: 15px;
    border-top: 1px solid rgb(var(--bg-2));
    gap: 15px;

    button {
        width: 120px;
    }
}
</style>
