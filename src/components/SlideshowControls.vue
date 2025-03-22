<script setup lang="ts">
import { setFont, setMonoFont } from '@/composables/useSettings'
import { friendlyTheme } from '@/composables/utils'
import { useSettingsStore } from '@/stores'
import type { Slide } from '@/types'
import {
    ChevronLeftIcon,
    ChevronRightIcon,
    PinIcon,
    PinOffIcon,
    SettingsIcon,
} from '@proicons/vue'
import { onClickOutside } from '@vueuse/core'
import { bundledThemes } from 'shiki'
import { computed, ref, watch } from 'vue'
import ContextMenu from './ContextMenu.vue'

const { slides } = defineProps<{ slides: Slide[] }>()
const currentSlide = defineModel<number>()

const slideList = slides.map(({ title }, i) => ({
    label: `#${i + 1}: ${title}`,
    action: () => (currentSlide.value = i),
}))

const s = useSettingsStore()

const menuIsPinned = ref(JSON.parse(localStorage.getItem('menuIsPinned')) ?? false)
watch(menuIsPinned, v => localStorage.setItem('menuIsPinned', JSON.stringify(v)))
const pinLabel = computed(() => (menuIsPinned.value === true ? 'Unpin' : 'Pin'))

const settingsItems = [
    {
        label: computed(() => `Theme: ${friendlyTheme(s.theme || 'Default')}`),
        action: s.showThemeSelection,
    },
    {
        label: computed(() => `Font: ${s.font || 'Default'}`),
        action: setFont,
    },
    {
        label: computed(() => `Monospace Font: ${s.monoFont || 'Default'}`),
        action: setMonoFont,
    },
    { label: 'separator' },
    {
        label: 'Go to First Slide',
        action: () => (currentSlide.value = 0),
    },
    { label: 'separator' },
    {
        label: 'Reset Settings',
        action: () =>
            confirm('Are you sure you want to reset all settings?') &&
            localStorage.clear(),
    },
]

const themes = Object.keys(bundledThemes)
const themePickerItems = [
    { label: 'Theme' },
    { label: 'separator' },
    ...themes.map(theme => ({
        label: friendlyTheme(theme),
        action: () => s.setTheme(theme),
    })),
]
const themeMenu = ref()
onClickOutside(themeMenu, s.hideThemeSelection)
</script>
<template>
    <div :class="['SlideshowControls', { pinned: menuIsPinned }]">
        <ContextMenu
            v-show="s.isSelectingTheme"
            ref="themeMenu"
            class="ThemeMenu"
            :items="themePickerItems" />
        <button
            class="slideAction FlyoutOpener"
            tabindex="0"
            :title="pinLabel"
            :aria-label="pinLabel"
            @click="() => (menuIsPinned = !menuIsPinned)">
            <PinOffIcon v-if="menuIsPinned" />
            <PinIcon v-else />
        </button>
        <div class="slideViewerShortcut">
            <button
                class="slideArrow"
                style="border-radius: 100px 0 0 100px; padding-left: 6px"
                title="Previous Slide"
                aria-label="Previous Slide"
                @click="() => currentSlide !== 0 && currentSlide--">
                <ChevronLeftIcon :size="20" />
            </button>
            <button
                class="slideViewerBtn FlyoutOpener"
                tabindex="0"
                title="Show All Slides"
                aria-label="Show All Slides">
                <span>Slide {{ currentSlide + 1 }} of {{ slides.length }}</span>
                <ContextMenu :items="slideList" />
            </button>
            <button
                class="slideArrow"
                title="Next Slide"
                aria-label="Next Slide"
                style="border-radius: 0 100px 100px 0; padding-right: 6px"
                @click="() => currentSlide < slides.length - 1 && currentSlide++">
                <ChevronRightIcon :size="20" />
            </button>
        </div>
        <button
            class="slideAction FlyoutOpener"
            tabindex="0"
            title="Settings"
            aria-label="Open Settings">
            <SettingsIcon />
            <ContextMenu :items="settingsItems" />
        </button>
    </div>
</template>
<style scoped lang="scss">
.SlideshowControls {
    display: flex;
    position: fixed;
    left: 50%;
    bottom: 3%;
    transform: translateX(-50%);
    padding: 10px;
    background: rgba(var(--bg-0), 0.9);
    backdrop-filter: blur(10px) brightness(1.5);
    border: 1px solid rgb(var(--bg-2), 0.5);
    box-shadow: var(--shadow-med);
    border-radius: 100px;
    transition: opacity 0.2s;
    gap: 4px;
    z-index: 4;

    &:not(.pinned) {
        opacity: 0;
    }

    &:hover,
    &:focus-within {
        opacity: 1;
    }

    .slideAction {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 6px;
        background: none;
        transition: background 0.15s;
        border-radius: 100px;

        &:hover {
            background: rgb(var(--bg-1));
        }
    }
}

.slideViewerShortcut {
    padding: 0;
    font-weight: 500;
    display: flex;
    align-items: center;
    background: rgb(var(--bg-1));
    border-radius: 100px;

    .slideArrow,
    .slideViewerBtn {
        background: none;
        transition: background 0.15s;
        display: flex;
        height: 100%;
        align-items: center;
        justify-content: center;
        text-align: center;

        &:hover {
            background: rgb(var(--bg-2));
        }
    }

    .slideArrow {
        padding: 3px;
        width: 35px;
        display: flex;
        position: relative;
    }

    .slideViewerBtn {
        padding: 5px 8px;
        border-radius: 0;
        border-inline: 2px solid rgb(var(--bg-2));
    }
}

.ThemeMenu {
    display: flex;
    min-width: 250px;

    & > :deep(button:first-child) {
        opacity: 0.6;
        cursor: default;
        font-weight: 500;
        background: none !important;
    }
}

.slideViewerShortcut :deep(.ContextMenu) {
    min-width: 300px;
}
</style>