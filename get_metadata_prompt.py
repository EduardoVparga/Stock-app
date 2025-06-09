import os

files_to_seach = [
    "src/components/stock-actions/columns.ts",
    "src/components/stock-actions/DateRangePicker.vue",
    "src/components/stock-actions/StockActionFilters.vue",
    "src/components/stock-actions/StockActionsClient.vue",
    "src/components/stock-actions/StockActionsTable.vue",
    "src/components/stock-actions/StockDetailModal.vue",
    
    "src/components/ui/Badge.vue",
    "src/components/ui/Button.vue",
    "src/components/ui/Input.vue",
    "src/components/ui/Select.vue",
    "src/components/ui/SelectContent.vue",
    "src/components/ui/SelectItem.vue",
    "src/components/ui/SelectParts.vue",
    "src/components/ui/SelectTrigger.vue",
    "src/components/ui/SelectValue.vue",
    "src/components/ui/Separator.vue",
    "src/components/ui/Sheet.vue",
    "src/components/ui/SheetContent.vue",
    "src/components/ui/SheetDescription.vue",
    "src/components/ui/SheetFooter.vue",
    "src/components/ui/SheetHeader.vue",
    "src/components/ui/SheetTitle.vue",
    "src/components/ui/Table.vue",
    "src/components/ui/TableBody.vue",
    "src/components/ui/TableCell.vue",
    "src/components/ui/TableHead.vue",
    "src/components/ui/TableHeader.vue",
    "src/components/ui/TableParts.vue",
    "src/components/ui/TableRow.vue",

    "src/lib/utils.ts",
    "src/stores/stockStore.ts",
    "src/types/stock-action.ts",
    "src/views/HomePage.vue",
    "src/App.vue",
    "src/main.ts",
    "src/router.ts",

    "package.json",
]

# Check if the files exist in the current directory
for file in files_to_seach:
    if os.path.exists(file):
        print(f"File found: {file}")
    else:
        print(f"File not found: {file}")

# get text from the files
def get_text_from_files(files):
    text = ""
    for file in files:
        if os.path.exists(file):
            print(f"Reading file: {file}")
            with open(file, 'r', encoding="utf-8") as f:
                text += f"# file name: {file} ------------------------------------------\n{f.read()}\n\n"
    return text

# generate .txt with the text from the files
text = get_text_from_files(files_to_seach)
with open("metadata.txt", 'w', encoding="utf-8") as f:
    f.write(text)
