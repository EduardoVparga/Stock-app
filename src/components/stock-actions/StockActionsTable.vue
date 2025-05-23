<template>
    <div class="w-full">
      <div class="rounded-md border shadow-sm bg-card">
        <Table>
          <TableHeader>
            <TableRow v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
              <TableHead
                v-for="header in headerGroup.headers"
                :key="header.id"
                :colSpan="header.colSpan"
                class="whitespace-nowrap px-3 py-3 text-sm"
                :style="{ width: header.getSize() !== 150 ? `${header.getSize()}px` : undefined }"
              >
                <FlexRender
                  v-if="!header.isPlaceholder"
                  :render="header.column.columnDef.header"
                  :props="header.getContext()"
                />
              </TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            <template v-if="table.getRowModel().rows?.length">
              <TableRow
                v-for="row in table.getRowModel().rows"
                :key="row.id"
                :data-state="row.getIsSelected() ? 'selected' : undefined"
                @click="() => onRowClick(row.original as TData)"
                class="cursor-pointer hover:bg-muted/50 transition-colors"
              >
                <TableCell
                  v-for="cell in row.getVisibleCells()"
                  :key="cell.id"
                  class="px-3 py-3 text-sm"
                >
                  <FlexRender
                    :render="cell.column.columnDef.cell"
                    :props="cell.getContext()"
                  />
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
      <div class="flex items-center justify-end space-x-2 py-4">
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
  </template>
  
  <script setup lang="ts" generic="TData extends StockAction, TValue">
  import { ref, onMounted, onBeforeUnmount, watchEffect, toRefs } from 'vue'
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
  
  const hasMounted = ref(false);
  
  onMounted(() => {
    hasMounted.value = true;
  });
  
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
      if (window.innerWidth < 768) { // Tailwind's 'md' breakpoint
        table.setColumnVisibility({
          ...columnVisibility.value, // preserve other explicit visibilities
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
      handleResize(); // Initial check
      window.addEventListener('resize', handleResize);
    }
  });
  
  onBeforeUnmount(() => {
    if (typeof window !== 'undefined') {
      window.removeEventListener('resize', handleResize);
    }
  });
  
  // The hasMounted check for rendering is less critical in Vue for this specific resize logic,
  // as DOM access is typically safer in onMounted. If there were truly conditional rendering
  // based on window that could cause hydration mismatch, techniques like <ClientOnly> or
  // v-if="hasMounted" on the whole component might be used. For now, the logic inside
  // onMounted should suffice.
  </script>