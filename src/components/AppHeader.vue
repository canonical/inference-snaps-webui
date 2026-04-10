<script setup lang="ts">
import { useChatStore } from '@/stores/chat'

const store = useChatStore()
</script>

<template>
  <header class="p-navigation" role="banner">
    <div class="p-navigation__row">
      <div class="p-navigation__banner">
        <div class="p-navigation__tagged-logo">
          <a class="p-navigation__link" href="#">
            <div class="p-navigation__logo-tag">
              <img
                src="https://assets.ubuntu.com/v1/82818827-CoF_white.svg"
                alt="Canonical"
                width="32"
                height="32"
              />
            </div>
            <span class="p-navigation__logo-title">{{ store.uiTitle }}</span>
          </a>
        </div>
      </div>
      <button
        class="settings-toggle"
        :aria-expanded="store.isSettingsPanelOpen"
        :aria-label="store.isSettingsPanelOpen ? 'Hide settings' : 'Show settings'"
        title="Settings"
        @click="store.isSettingsPanelOpen = !store.isSettingsPanelOpen"
      >
        <i class="p-icon--information"></i>
      </button>
    </div>
  </header>
</template>

<style scoped>
.p-navigation__logo-title {
  padding-left: 0.75rem;
}

.settings-toggle {
  position: absolute;
  top: 50%;
  right: 1rem;
  transform: translateY(-50%);
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.25rem 0.5rem;
  border-radius: 0.25rem;
  opacity: 0.35;
  transition: opacity 0.15s;
  line-height: 0;
}

.settings-toggle:hover {
  opacity: 0.75;
}

.settings-toggle .p-icon--information {
  display: inline-block;
  width: 1.5rem;
  height: 1.5rem;
  filter: brightness(0) invert(1);
}
</style>

