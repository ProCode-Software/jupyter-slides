<script setup lang="ts">
import type { Slide } from '@/types'
import { ChevronLeftIcon, ChevronRightIcon, SettingsIcon } from '@proicons/vue'
import ContextMenu from './ContextMenu.vue'
import { setFont, setMonoFont } from '@/composables/useSettings';
import { useSettingsStore } from '@/stores';
import { computed } from 'vue';
import { friendlyTheme } from '@/composables/utils';

const { slides } = defineProps<{ slides: Slide[] }>()
const currentSlide = defineModel<number>()

const s = useSettingsStore()

const settingsItems = [
    {
        label: computed(() => `Theme: ${friendlyTheme(s.theme || 'Default')}`),
        action: s.showThemeSelection
    },
    {
        label: computed(() => `Font: ${s.font || 'Default'}`),
        action: setFont
    },
    {
        label: computed(() => `Monospace Font: ${s.monoFont || 'Default'}`),
        action: setMonoFont
    },
    { label: 'separator' },
    {
        label: 'Go to First Slide',
        action: () => currentSlide.value = 0
    },
    { label: 'separator' },
    {
        label: 'Reset Settings',
        action: () => prompt('Are you sure you want to reset all settings?')
            && localStorage.clear()
    }
]
</script>
<template>
    <div class="SlideshowControls">
        <div class="slideViewerShortcut">
            <button class="slideArrow"
                @click="() => currentSlide !== 0 && currentSlide--">
                <ChevronLeftIcon :size="20" />
            </button>
            <button class="slideViewerBtn">
                Slide {{ currentSlide + 1 }} of
                {{ slides.length }}
            </button>
            <button class="slideArrow"
                @click="() => currentSlide <= slides.length - 1 && currentSlide++">
                <ChevronRightIcon :size="20" />
            </button>
        </div>
        <button class="slideAction FlyoutOpener">
            <SettingsIcon />
            <ContextMenu :items="settingsItems" />
        </button>
    </div>
</template>
<style scoped lang="scss">
.SlideshowControls {
    display: flex;
    position: absolute;
    left: 50%;
    bottom: 3%;
    transform: translateX(-50%);
    padding: 10px;
    background: rgba(var(--bg-0), 0.75);
    backdrop-filter: blur(10px);
    border: 1px solid rgb(var(--bg-2), 0.5);
    box-shadow: var(--shadow-med);
    border-radius: 100px;
    opacity: 0;
    transition: opacity 0.2s;
    gap: 8px;

    &:hover {
        opacity: 1;
    }

    .slideAction {
        display: flex;
        align-items: center;
        justify-content: center;
        padding: 6px;
        background: none;
        transition: 0.15s;
        border-radius: 100px;

        &:hover {
            background: rgb(var(--bg-1));
        }
    }

    .slideViewerShortcut {
        padding: 4px;
        font-weight: 500;
        display: flex;
        align-items: center;
        background: rgb(var(--bg-1));
        border-radius: 100px;

        .slideArrow,
        .slideViewerBtn {
            border-radius: 100px;
            background: none;
            transition: 0.15s;
            display: flex;
            height: 30px;

            &:hover {
                background: rgb(var(--bg-2));
            }
        }

        .slideArrow {
            padding: 5px;
            width: 30px;
            display: flex;
        }

        .slideViewerBtn {
            padding: 5px 8px;
        }
    }
}
</style>
