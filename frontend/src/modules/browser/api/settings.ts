import type { BrowserSettings } from '../types'
import { createDefaultBrowserSettings, ensureMockFallbackAllowed, getBindings } from './runtime'

export async function fetchBrowserSettings(): Promise<BrowserSettings> {
  const bindings: any = await getBindings()
  if (bindings?.GetBrowserSettings) {
    return (await bindings.GetBrowserSettings()) || createDefaultBrowserSettings()
  }
  ensureMockFallbackAllowed()
  return createDefaultBrowserSettings()
}

export async function saveBrowserSettings(settings: BrowserSettings): Promise<boolean> {
  const bindings: any = await getBindings()
  if (bindings?.SaveBrowserSettings) {
    await bindings.SaveBrowserSettings(settings)
    return true
  }
  ensureMockFallbackAllowed()
  return true
}
