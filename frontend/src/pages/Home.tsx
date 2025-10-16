import { useState } from "react"
import MainLayout from "../components/layouts/MainLayout"

export default function Home() {
  const [count, setCount] = useState(0)

  return (
    <MainLayout>
      <div className="min-h-screen bg-gradient-to-br from-slate-900 via-slate-800 to-slate-900 text-slate-100 antialiased">
        <div className="container mx-auto px-4 py-16">
          <div className="mx-auto max-w-xl text-center">
            <h1 className="text-4xl md:text-5xl font-extrabold tracking-tight bg-clip-text text-transparent bg-gradient-to-r from-emerald-400 to-cyan-400">
              cuhara.qua
            </h1>
            <p className="mt-3 text-slate-300">
              Vite + React + Tailwind starter
            </p>

            <div className="mt-8 rounded-2xl bg-slate-800/60 border border-slate-700 shadow-xl p-8 backdrop-blur">
              <p className="text-slate-300">Sayaç</p>
              <div className="mt-3 flex items-center justify-center gap-4">
                <button
                  onClick={() => setCount((c) => c - 1)}
                  className="inline-flex items-center rounded-lg bg-slate-700 hover:bg-slate-600 active:bg-slate-700/80 px-4 py-2 text-sm font-medium text-slate-100 transition"
                >
                  -1
                </button>
                <span className="text-3xl font-bold tabular-nums">{count}</span>
                <button
                  onClick={() => setCount((c) => c + 1)}
                  className="inline-flex items-center rounded-lg bg-emerald-500 hover:bg-emerald-400 active:bg-emerald-500/80 px-4 py-2 text-sm font-semibold text-emerald-950 transition"
                >
                  +1
                </button>
              </div>

              <p className="mt-6 text-sm text-slate-400">
                Edit <code className="rounded bg-slate-900/50 px-1 py-0.5">src/App.tsx</code> and save to test HMR.
              </p>
            </div>

            <div className="mt-10 flex items-center justify-center gap-3 text-sm text-slate-400">
              <a className="hover:text-slate-200 transition underline underline-offset-4" href="https://vite.dev" target="_blank" rel="noreferrer">
                Vite
              </a>
              <span>•</span>
              <a className="hover:text-slate-200 transition underline underline-offset-4" href="https://react.dev" target="_blank" rel="noreferrer">
                React
              </a>
              <span>•</span>
              <a className="hover:text-slate-200 transition underline underline-offset-4" href="https://tailwindcss.com" target="_blank" rel="noreferrer">
                Tailwind
              </a>
            </div>
          </div>
        </div>
      </div>
    </MainLayout>
  )
}