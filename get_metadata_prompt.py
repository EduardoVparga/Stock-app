import os

files_to_seach = [
    # "src/App.vue",
    # "src/assets/main.css",
    # "src/assets/vue.svg",
    "src/components/stock-actions/DateRangePicker.vue",
    "src/components/stock-actions/RankedStockCard.vue",
    "src/components/stock-actions/StockActionFilters.vue",
    "src/components/stock-actions/StockActionsClient.vue",
    "src/components/stock-actions/StockActionsTable.vue",
    "src/components/stock-actions/StockDetailModal.vue",
    "src/components/stock-actions/TopRankedStocks.vue",
    "src/components/stock-actions/columns.ts",
    "src/components/ui/Badge.vue",
    "src/components/ui/Button.vue",
    "src/components/ui/Input.vue",
    "src/components/ui/Select.vue",
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
    "src/main.ts",
    # "src/router.ts",
    "src/stores/stockStore.ts",
    "src/types/stock-action.ts",
    "src/views/HomePage.vue",
    # "src/vite-env.d.ts",

    "package.json",

    "tailwind.config.js",
    "src/assets/main.css"
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



# from pathlib import Path
# from typing import List

# def scan_files(root_dir: str) -> List[str]:
#     """
#     Escanea recursivamente todos los archivos bajo root_dir
#     y devuelve una lista de rutas relativas.
    
#     :param root_dir: Ruta al directorio raíz a escanear.
#     :return: Lista de rutas de archivos (strings), relativas a root_dir.
#     """
#     root = Path(root_dir)
#     if not root.is_dir():
#         raise ValueError(f"El path proporcionado no es un directorio válido: {root_dir}")

#     files = [
#         str(path.relative_to(root))
#         for path in root.rglob('*')
#         if path.is_file()
#     ]
#     return sorted(files)


# # Ejemplo de uso:
# if __name__ == "__main__":
#     files_to_search = scan_files("src")
#     for f in files_to_search:
#         print(f"src/{f.replace("\\", "/")}")