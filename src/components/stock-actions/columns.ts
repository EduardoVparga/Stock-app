import { h } from 'vue'
import type { ColumnDef } from "@tanstack/vue-table"
import { format } from "date-fns"
import type { StockAction } from "@/types/stock-action"
import Badge from "@/components/ui/Badge.vue"
import { formatPrice } from '@/lib/utils'

export const columns: ColumnDef<StockAction>[] = [
  {
    accessorKey: "company",
    header: "Company",
    cell: ({ row }) => h('div', { class: 'font-medium' }, row.getValue("company")),
  },
  {
    accessorKey: "ticker",
    header: "Ticker",
    cell: ({ row }) => h(Badge, { variant: "secondary" }, () => row.getValue("ticker")),
  },
  {
    accessorKey: "brokerage",
    header: "Brokerage",
    cell: ({ row }) => h('div', {}, row.getValue("brokerage")),
  },
  {
    accessorKey: "action",
    header: "Rating",
    cell: ({ row }) => {
      const action = row.getValue("action")?.toLowerCase() ?? '';
      
      let variant: 'default' | 'destructive' | 'outline' = 'outline'; 
      let text = 'Hold';

      if (action.includes('upgrade') || action.includes('initiated')) {
        variant = 'default'; 
        text = 'Buy';
      } else if (action.includes('downgrade')) {
        variant = 'destructive'; 
        text = 'Sell';
      }
      
      return h(Badge, { variant, class: 'capitalize px-3 py-1 text-xs' }, () => text)
    },
  },
  {
    accessorKey: "spot",
    header: "Price",
    cell: ({ row }) => h('div', { class: 'font-mono' }, formatPrice(row.getValue("spot"))),
  },
  {
    id: "change",
    header: "Change",
    cell: ({ row }) => {
        const from = row.original.target_from;
        const to = row.original.target_to;
        if (to === null || from === null) return h('span', { class: 'text-muted-foreground' }, 'N/A');
        
        const change = to - from;
        const isPositive = change >= 0;
        const colorClass = isPositive ? 'text-green-500' : 'text-red-500';
        const formattedChange = `${isPositive ? '+' : ''}${formatPrice(change)}`;
        
        return h('div', { class: `${colorClass} font-mono` }, formattedChange);
    }
  },
  {
    accessorKey: "time",
    header: "Date",
    cell: ({ row }) => {
      const date = new Date(row.getValue("time"));
      return h('div', { class: 'text-muted-foreground' }, format(date, "MMM dd, yyyy"));
    },
  },
]