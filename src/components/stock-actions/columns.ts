// src/components/stock-actions/columns.ts
import { h } from 'vue'
import type { ColumnDef } from "@tanstack/vue-table"
import { format } from "date-fns"
import { ArrowUpDown } from "lucide-vue-next"

import type { StockAction } from "@/types/stock-action"
import Button from "@/components/ui/Button.vue" // Vue Button
import Badge from "@/components/ui/Badge.vue"   // Vue Badge
import { formatPrice } from "@/lib/utils"

export const columns: ColumnDef<StockAction>[] = [
  {
    accessorKey: "time",
    header: ({ column }) => {
      return h(Button, {
        variant: "ghost",
        onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }, () => [
        'Date',
        h(ArrowUpDown, { class: "ml-2 h-4 w-4" })
      ])
    },
    cell: ({ row }) => {
      const date = new Date(row.getValue("time"))
      return h('div', { class: 'text-sm' }, format(date, "MMM dd, yyyy"))
    },
    // Enable sorting for this column if desired
    enableSorting: true,
  },
  {
    accessorKey: "brokerage",
    header: ({ column }) => {
      return h(Button, {
        variant: "ghost",
        onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }, () => [
        'Brokerage',
        h(ArrowUpDown, { class: "ml-2 h-4 w-4" })
      ])
    },
    cell: ({ row }) => h('div', { class: 'text-sm' }, row.getValue("brokerage")),
    enableSorting: true,
  },
  {
    accessorKey: "ticker",
    header: ({ column }) => {
      return h(Button, {
        variant: "ghost",
        onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }, () => [
        'Ticker',
        h(ArrowUpDown, { class: "ml-2 h-4 w-4" })
      ])
    },
    cell: ({ row }) => h(Badge, { variant: "secondary" }, () => row.getValue("ticker")),
    enableSorting: true,
  },
  {
    accessorKey: "company",
    header: ({ column }) => {
      return h(Button, {
        variant: "ghost",
        onClick: () => column.toggleSorting(column.getIsSorted() === "asc"),
      }, () => [
        'Company',
        h(ArrowUpDown, { class: "ml-2 h-4 w-4" })
      ])
    },
    cell: ({ row }) => h('div', { class: 'font-medium text-sm' }, row.getValue("company")),
    enableSorting: true,
  },
  {
    accessorKey: "action",
    header: "Action",
    cell: ({ row }) => h('div', { class: 'capitalize text-sm' }, row.getValue("action")),
  },
  {
    id: "ratingChange",
    header: "Rating",
    cell: ({ row }) => {
      const ratingFrom = row.original.rating_from || "N/A"
      const ratingTo = row.original.rating_to
      return h('div', { class: 'text-sm whitespace-nowrap' }, `${ratingFrom} → ${ratingTo}`)
    },
  },
  {
    id: "priceTargetChange",
    header: "Price Target",
    cell: ({ row }) => {
      const targetFrom = row.original.target_from
      const targetTo = row.original.target_to

      const formattedFrom = formatPrice(targetFrom)
      const formattedTo = formatPrice(targetTo)
      
      let colorClass = "text-foreground" // Default color
      if (targetTo !== null && targetFrom !== null) {
        if (targetTo > targetFrom) {
          colorClass = "text-green-600 dark:text-green-500"
        } else if (targetTo < targetFrom) {
          colorClass = "text-red-600 dark:text-red-500"
        }
      } else if (targetTo !== null && targetFrom === null && row.original.action.toLowerCase().includes("initiated")) {
        colorClass = "text-green-600 dark:text-green-500"
      }

      return h('div', { class: `text-sm font-semibold whitespace-nowrap ${colorClass}` }, 
        `${formattedFrom} → ${formattedTo}`
      )
    },
  },
]