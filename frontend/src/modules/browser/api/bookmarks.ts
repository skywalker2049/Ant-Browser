import type { BookmarkSyncResult, BrowserBookmark } from '../types'
import { ensureMockFallbackAllowed, getBindings } from './runtime'

export async function fetchBookmarks(): Promise<BrowserBookmark[]> {
  const bindings: any = await getBindings()
  if (bindings?.BookmarkList) {
    return (await bindings.BookmarkList()) || []
  }
  ensureMockFallbackAllowed()
  return [
    { name: 'Google', url: 'https://www.google.com/', openOnStart: false },
    { name: 'Gmail', url: 'https://mail.google.com/', openOnStart: false },
    { name: 'Claude', url: 'https://claude.ai/', openOnStart: false },
    { name: 'ChatGPT', url: 'https://chatgpt.com/', openOnStart: false },
    { name: 'YouTube', url: 'https://www.youtube.com/', openOnStart: false },
  ]
}

export async function saveBookmarks(items: BrowserBookmark[]): Promise<boolean> {
  const bindings: any = await getBindings()
  if (bindings?.BookmarkSave) {
    await bindings.BookmarkSave(items)
    return true
  }
  ensureMockFallbackAllowed()
  return true
}

export async function resetBookmarks(): Promise<boolean> {
  const bindings: any = await getBindings()
  if (bindings?.BookmarkReset) {
    await bindings.BookmarkReset()
    return true
  }
  ensureMockFallbackAllowed()
  return true
}

export async function syncBookmarksToProfiles(): Promise<BookmarkSyncResult> {
  const bindings: any = await getBindings()
  if (bindings?.BookmarkSyncToProfiles) {
    return await bindings.BookmarkSyncToProfiles()
  }
  ensureMockFallbackAllowed()
  return {
    total: 0,
    synced: 0,
    skipped: 0,
    failed: 0,
    skippedList: [],
    failedList: [],
  }
}
