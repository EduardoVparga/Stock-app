<template>
  <div class="w-full">
    <div class="rounded-lg border border-border bg-card">
      <Table>
        <TableHeader>
          <TableRow class="border-b border-border hover:bg-transparent">
            <TableHead
              v-for="header in table.getHeaderGroups()[0].headers"
              :key="header.id"
              class="text-muted-foreground font-semibold uppercase text-xs tracking-wider"
            >
              <FlexRender :render="header.column.columnDef.header" :props="header.getContext()" />
            </TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          <template v-if="table.getRowModel().rows?.length">
            <TableRow
              v-for="row in table.getRowModel().rows"
              :key="row.id"
              @click="() => onRowClick(row.original as TData)"
              class="border-b border-border/50 hover:bg-secondary/50 cursor-pointer"
            >
              <TableCell v-for="cell in row.getVisibleCells()" :key="cell.id" class="py-3 px-4">
                <FlexRender :render="cell.column.columnDef.cell" :props="cell.getContext()" />
              </TableCell>
            </TableRow>
          </template>
          <template v-else>
            <TableRow>
              <TableCell :colSpan="columns.length" class="h-24 text-center">
                No results found.
              </TableCell>
            </TableRow>
          </template>
        </TableBody>
      </Table>
    </div>

    <div class="flex items-center justify-between space-x-2 py-4 mt-4">
      <span class="text-sm text-muted-foreground">
        Page {{ table.getState().pagination.pageIndex + 1 }} of {{ table.getPageCount() }}
      </span>
      <div class="flex items-center space-x-2">
        <Button 
            variant="outline" 
            size="sm" 
            @click="table.previousPage()" 
            :disabled="!table.getCanPreviousPage()"
        >
          Previous
        </Button>
        <Button 
            variant="outline" 
            size="sm" 
            @click="table.nextPage()" 
            :disabled="!table.getCanNextPage()"
        >
          Next
        </Button>
      </div>
    </div>
  </div>
</template>
  
<script setup lang="ts" generic="TData extends StockAction, TValue">
  import { ref, onMounted, onBeforeUnmount, toRefs } from 'vue'
  import type {
    ColumnDef,
    SortingState,
    VisibilityState,
    ColumnFiltersState,
  } from "@tanstack/vue-table"
  import {
    FlexRender,
    getCoreRowModel,
    getFilteredRowModel,
    getPaginationRowModel,
    getSortedRowModel,
    useVueTable,
  } from "@tanstack/vue-table"
  
  import Table from "@/components/ui/Table.vue"
  import TableBody from "@/components/ui/TableBody.vue"
  import TableCell from "@/components/ui/TableCell.vue"
  import TableHead from "@/components/ui/TableHead.vue"
  import TableHeader from "@/components/ui/TableHeader.vue"
  import TableRow from "@/components/ui/TableRow.vue"
  import Button from "@/components/ui/Button.vue"
  import type { StockAction } from "@/types/stock-action"
  
  interface DataTableProps {
    columns: ColumnDef<TData, TValue>[]
    data: TData[]
    onRowClick: (row: TData) => void
  }
  
  const props = defineProps<DataTableProps>()
  const { columns, data, onRowClick } = toRefs(props)
  
  const sorting = ref<SortingState>([])
  const columnFilters = ref<ColumnFiltersState>([])
  const columnVisibility = ref<VisibilityState>({
    'brokerage': false, 
    'ratingChange': false,
  })
    
  const table = useVueTable({
    get data() { return data.value },
    get columns() { return columns.value },
    state: {
      get sorting() { return sorting.value },
      get columnFilters() { return columnFilters.value },
      get columnVisibility() { return columnVisibility.value },
    },
    onSortingChange: (updater) => {
      sorting.value = typeof updater === 'function' ? updater(sorting.value) : updater
    },
    onColumnFiltersChange: (updater) => {
      columnFilters.value = typeof updater === 'function' ? updater(columnFilters.value) : updater
    },
    onColumnVisibilityChange: (updater) => {
      columnVisibility.value = typeof updater === 'function' ? updater(columnVisibility.value) : updater
    },
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    initialState: {
      pagination: {
        pageSize: 10,
      },
    },
  })
  
  const handleResize = () => {
    if (typeof window !== 'undefined') {
      if (window.innerWidth < 768) { 
        table.setColumnVisibility({
          ...columnVisibility.value,
          'brokerage': false,
          'ratingChange': false,
          'time': false,
        });
      } else {
        table.setColumnVisibility({
          ...columnVisibility.value,
          'brokerage': true,
          'ratingChange': true,
          'time': true,
        });
      }
    }
  };
  
  onMounted(() => {
    if (typeof window !== 'undefined') {
      handleResize(); 
      window.addEventListener('resize', handleResize);
    }
  });
  
  onBeforeUnmount(() => {
    if (typeof window !== 'undefined') {
      window.removeEventListener('resize', handleResize);
    }
  });
</script>