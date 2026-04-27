<script setup lang="ts">
import { marked } from 'marked'
import DOMPurify from 'dompurify'
import { useChatStore } from '@/stores/chat'
import type { ChatMessage, MessageContentPart } from '@/types'

const props = defineProps<{
  message: ChatMessage
  index: number
}>()

const emit = defineEmits<{
  retry: [index: number]
}>()

const store = useChatStore()

function formatPlain(text: string | null | undefined): string {
  if (!text) return ''
  return DOMPurify.sanitize(
    String(text)
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/"/g, '&quot;')
      .replace(/'/g, '&#039;')
      .replace(/\n/g, '<br>'),
  )
}

function formatMarkdown(text: string | null | undefined): string {
  if (!text) return ''
  return DOMPurify.sanitize(marked.parse(String(text)) as string)
}

function formatText(text: string | null | undefined): string {
  return store.isMarkdown ? formatMarkdown(text) : formatPlain(text)
}

function getUserText(): string {
  const c = props.message.content
  if (Array.isArray(c)) {
    return (c as MessageContentPart[]).find((p) => p.type === 'text')?.text ?? ''
  }
  return c
}

function toggleReasoning() {
  // Mutate via the store's messages array to avoid prop mutation lint error
  const storeMsg = store.messages[props.index]
  if (storeMsg) storeMsg.reasoningOpen = !storeMsg.reasoningOpen
}
</script>

<template>
  <!-- User message -->
  <div v-if="message.role === 'user'" class="chat-message chat-message--user">
    <div class="chat-message__avatar" aria-label="You">You</div>
    <div class="chat-message__bubble">
      <div
        class="chat-message__text"
        :class="{ 'chat-message__text--markdown': store.isMarkdown }"
        v-html="formatText(getUserText())"
      ></div>
      <div v-if="message.images && message.images.length" class="chat-message__images">
        <img
          v-for="(img, i) in message.images"
          :key="i"
          :src="img"
          alt="Attached image"
          class="chat-message__image-preview"
          title="Click to view full size"
          @click="store.openLightbox(img)"
        />
      </div>
      <div v-if="message.cancelled" class="message-cancelled-notice">Cancelled</div>
    </div>
  </div>

  <!-- Assistant message -->
  <div v-else class="chat-message chat-message--assistant">
    <div class="chat-message__avatar" aria-label="Assistant">AI</div>
    <div class="chat-message__bubble">

      <!-- Reasoning section -->
      <div v-if="message.reasoning" class="reasoning-block">
        <button
          class="reasoning-toggle"
          :aria-expanded="message.reasoningOpen"
          @click="toggleReasoning"
        >
          <span class="reasoning-toggle__icon">{{ message.reasoningOpen ? '▾' : '▸' }}</span>
          <span>{{ message.isStreaming ? 'Thinking…' : 'View reasoning' }}</span>
          <span v-if="message.isStreaming" class="reasoning-pulse"></span>
        </button>
        <div v-show="message.reasoningOpen" class="reasoning-content" :class="{ 'reasoning-content--markdown': store.isMarkdown }">
          <div v-html="formatText(message.reasoning)"></div>
        </div>
      </div>

      <!-- Error message -->
      <div v-if="message.friendlyError" class="chat-error">
        <i class="p-icon--error chat-error__icon"></i>
        <span class="chat-error__text">{{ message.friendlyError }}</span>
      </div>

      <!-- Response text -->
      <div v-else class="chat-message__text" :class="{ 'chat-message__text--markdown': store.isMarkdown }">
        <span v-html="formatText(typeof message.content === 'string' ? message.content : '')"></span>
        <span v-if="message.isStreaming && message.content" class="stream-cursor"></span>
      </div>

      <!-- Retry / details buttons (shown on error) -->
      <div v-if="message.friendlyError" class="chat-error__actions">
        <button
          class="p-button--neutral is-small u-no-margin--bottom"
          @click="emit('retry', index)"
        >
          Retry
        </button>
        <button
          class="p-button--neutral is-small u-no-margin--bottom"
          @click="store.openErrorModal(message.errorDetail ?? '')"
        >
          View details
        </button>
      </div>

      <!-- Cancelled indicator -->
      <div v-if="message.cancelled" class="message-cancelled-notice">Cancelled</div>

      <!-- Spinner while waiting for first token -->
      <div v-if="message.isStreaming && !message.content && !message.reasoning" class="chat-spinner">
        <div class="p-icon--spinner u-animation--spin"></div>
        <span class="u-text--muted">Working…</span>
      </div>

    </div>
  </div>
</template>


<style scoped>
.chat-message {
  display: flex;
  gap: 0.75rem;
  max-width: 100%;
  animation: fadeInUp 0.15s ease-out;
}

.chat-message--user {
  flex-direction: row-reverse;
}

.chat-message__avatar {
  flex-shrink: 0;
  width: 2rem;
  height: 2rem;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.65rem;
  font-weight: 700;
  text-transform: uppercase;
  color: #fff;
  background-color: #666;
}

.chat-message--user .chat-message__avatar {
  background-color: #06c;
}

.chat-message--assistant .chat-message__avatar {
  background-color: #e95420;
}

.chat-message__bubble {
  max-width: min(70%, 640px);
  padding: 0.6rem 0.9rem;
  border-radius: 0.5rem;
  line-height: 1.5;
  word-break: break-word;
}

.chat-message--user .chat-message__bubble {
  background-color: #06c;
  color: #fff;
  border-bottom-right-radius: 0.1rem;
}

.chat-message--assistant .chat-message__bubble {
  background-color: #fff;
  color: #111;
  border: 1px solid #d9d9d9;
  border-bottom-left-radius: 0.1rem;
}

.chat-message__text {
  white-space: pre-wrap;
}

.chat-message__text--markdown {
  white-space: normal;
}

/* Markdown typography — pierce v-html content via :deep() */
.chat-message__text--markdown :deep(p) {
  margin-top: 0;
  margin-bottom: 0.75rem;
}

.chat-message__text--markdown :deep(p:last-child) {
  margin-bottom: 0;
}

.chat-message__text--markdown :deep(h1),
.chat-message__text--markdown :deep(h2),
.chat-message__text--markdown :deep(h3),
.chat-message__text--markdown :deep(h4),
.chat-message__text--markdown :deep(h5),
.chat-message__text--markdown :deep(h6) {
  margin-top: 1rem;
  margin-bottom: 0.4rem;
  font-weight: 700;
  line-height: 1.3;
}

.chat-message__text--markdown :deep(ul),
.chat-message__text--markdown :deep(ol) {
  padding-left: 1.5rem;
  margin-bottom: 0.75rem;
}

.chat-message__text--markdown :deep(li) {
  margin-bottom: 0.2rem;
}

.chat-message__text--markdown :deep(code) {
  font-family: monospace;
  font-size: 0.875em;
  padding: 0.1em 0.3em;
  border-radius: 0.2rem;
  background-color: rgba(0, 0, 0, 0.08);
}

.chat-message__text--markdown :deep(pre) {
  background-color: rgba(0, 0, 0, 0.06);
  border-radius: 0.4rem;
  padding: 0.75rem 1rem;
  overflow-x: auto;
  margin-bottom: 0.75rem;
}

.chat-message__text--markdown :deep(pre code) {
  background: none;
  padding: 0;
  font-size: 0.85em;
}

.chat-message__text--markdown :deep(blockquote) {
  border-left: 3px solid #d9d9d9;
  margin: 0 0 0.75rem;
  padding-left: 1rem;
  color: #666;
}

.chat-message__text--markdown :deep(a) {
  color: #06c;
  text-decoration: underline;
}

.chat-message__text--markdown :deep(hr) {
  border: none;
  border-top: 1px solid #d9d9d9;
  margin: 0.75rem 0;
}

.chat-message__text--markdown :deep(table) {
  border-collapse: collapse;
  margin-bottom: 0.75rem;
  width: 100%;
}

.chat-message__text--markdown :deep(th),
.chat-message__text--markdown :deep(td) {
  border: 1px solid #d9d9d9;
  padding: 0.3rem 0.6rem;
  text-align: left;
}

/* Override code/link colours inside user (blue) bubbles */
.chat-message--user .chat-message__text--markdown :deep(code) {
  background-color: rgba(255, 255, 255, 0.2);
  color: #fff;
}

.chat-message--user .chat-message__text--markdown :deep(pre) {
  background-color: rgba(255, 255, 255, 0.15);
}

.chat-message--user .chat-message__text--markdown :deep(pre code) {
  color: #fff;
}

.chat-message--user .chat-message__text--markdown :deep(a) {
  color: #cce4ff;
}

.chat-message--user .chat-message__text--markdown :deep(blockquote) {
  border-left-color: rgba(255, 255, 255, 0.5);
  color: rgba(255, 255, 255, 0.8);
}

/* Reasoning content markdown overrides */
.reasoning-content--markdown {
  white-space: normal;
}

.reasoning-content--markdown :deep(p) {
  margin-top: 0;
  margin-bottom: 0.5rem;
}

.reasoning-content--markdown :deep(p:last-child) {
  margin-bottom: 0;
}

.reasoning-content--markdown :deep(code) {
  font-family: monospace;
  font-size: 0.85em;
  padding: 0.1em 0.25em;
  border-radius: 0.2rem;
  background-color: rgba(0, 0, 0, 0.07);
}

.chat-message__images {
  margin-top: 0.5rem;
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.chat-message__image-preview {
  max-width: 200px;
  max-height: 150px;
  object-fit: contain;
  border-radius: 0.25rem;
  border: 1px solid rgba(0, 0, 0, 0.1);
  cursor: zoom-in;
  transition: opacity 0.15s, transform 0.15s;
}

.chat-message__image-preview:hover {
  opacity: 0.9;
  transform: scale(1.02);
}

.chat-error {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  flex-wrap: wrap;
}

.chat-error__icon {
  display: inline-block;
  width: 1.25rem;
  height: 1.25rem;
  flex-shrink: 0;
}

.chat-error__text {
  flex: 1;
}

.chat-error__actions {
  margin-top: 0.5rem;
  display: flex;
  gap: 0.4rem;
  flex-wrap: wrap;
}

.message-cancelled-notice {
  margin-top: 0.4rem;
  font-size: 0.75rem;
  color: #c7a000;
  font-style: italic;
}

.stream-cursor {
  display: inline-block;
  width: 2px;
  height: 1em;
  background-color: currentColor;
  margin-left: 1px;
  vertical-align: text-bottom;
  animation: blink 0.8s step-end infinite;
}

.chat-spinner {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.p-icon--spinner {
  display: inline-block;
  width: 1.25rem;
  height: 1.25rem;
  border: 2px solid #d9d9d9;
  border-top-color: #e95420;
  border-radius: 50%;
  flex-shrink: 0;
}

/* Reasoning block */
.reasoning-block {
  margin-bottom: 0.6rem;
  border: 1px solid #c4a882;
  border-radius: 0.4rem;
  overflow: hidden;
  background-color: #fffbf4;
}

.reasoning-toggle {
  display: flex;
  align-items: center;
  gap: 0.4rem;
  width: 100%;
  padding: 0.4rem 0.6rem;
  margin-bottom: 0;
  background: #fef3e2;
  border: none;
  cursor: pointer;
  font-size: 0.8rem;
  font-weight: 600;
  color: #7a5200;
  text-align: left;
  transition: background-color 0.1s;
}

.reasoning-toggle:hover {
  background-color: #fde8c2;
}

.reasoning-toggle__icon {
  font-size: 0.7rem;
  line-height: 1;
}

.reasoning-pulse {
  display: inline-block;
  width: 7px;
  height: 7px;
  border-radius: 50%;
  background-color: #e95420;
  animation: pulse 1s ease-in-out infinite;
  margin-left: 0.25rem;
}

.reasoning-content {
  padding: 0.5rem 0.7rem;
  font-size: 0.82rem;
  color: #5a4000;
  line-height: 1.55;
  border-top: 1px solid #c4a882;
  max-height: 12rem;
  overflow-y: auto;
  white-space: pre-wrap;
  word-break: break-word;
}

.reasoning-content--markdown {
  white-space: normal;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(6px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes blink {
  50% {
    opacity: 0;
  }
}

@keyframes pulse {
  0%,
  100% {
    opacity: 1;
    transform: scale(1);
  }
  50% {
    opacity: 0.5;
    transform: scale(0.85);
  }
}

@media (max-width: 600px) {
  .chat-message__bubble {
    max-width: 88%;
  }
}
</style>

