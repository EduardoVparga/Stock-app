# Equity Pulse - Panel de Acciones Bursátiles



**Equity Pulse** es una aplicación web moderna y responsiva, diseñada para visualizar y analizar las últimas acciones y calificaciones del mercado de valores emitidas por diferentes firmas de corretaje. La interfaz, limpia e intuitiva, permite a los usuarios identificar rápidamente las empresas con mejor rendimiento ("Top Movers"), filtrar datos a través de una tabla interactiva y ver detalles específicos de cada acción.

## ✨ Características Principales

*   **Panel de "Top Movers"**: Muestra de forma destacada las 5 acciones con mejor ranking para una identificación rápida de oportunidades.
*   **Filtros Avanzados**: Permite a los usuarios filtrar los datos por:
    *   Nombre de la empresa o ticker.
    *   Tipo de acción (Upgrade, Downgrade, Initiated, etc.).
    *   Firma de corretaje.
    *   Rango de fechas personalizado.
*   **Tabla de Datos Interactiva**: Una tabla potente y paginada (construida con TanStack Table) que muestra todas las acciones. Es ordenable y se adapta a diferentes tamaños de pantalla.
*   **Modal de Detalles**: Al hacer clic en una acción (ya sea en las tarjetas o en la tabla), se abre un modal con información detallada, incluyendo cambios en el precio objetivo y enlaces directos a Yahoo Finance y Seeking Alpha.
*   **Diseño Moderno y Responsivo**: Construido con Tailwind CSS, la interfaz se ve genial y funciona a la perfección en dispositivos de escritorio y móviles.
*   **Gestión de Estado Centralizada**: Utiliza Pinia para una gestión de estado predecible y eficiente, manejando la carga de datos, estados de carga y errores.

## 🚀 Pila Tecnológica (Tech Stack)

Este proyecto está construido con un conjunto de herramientas modernas del ecosistema de Vue.js, enfocándose en el rendimiento, la escalabilidad y la experiencia del desarrollador.

*   **Framework Principal**: [Vue.js 3](https://vuejs.org/) (usando Composition API y `<script setup>`)
*   **Lenguaje**: [TypeScript](https://www.typescriptlang.org/)
*   **Estilos**: [Tailwind CSS](https://tailwindcss.com/)
    *   **Componentes de UI**: Sistema de componentes personalizado inspirado en [shadcn/ui](https://ui.shadcn.com/), utilizando `class-variance-authority` y `tailwind-merge`.
    *   **Iconos**: [Lucide Vue Next](https://lucide.dev/)
*   **Gestión de Estado**: [Pinia](https://pinia.vuejs.org/)
*   **Manejo de Datos y Tablas**: [TanStack Table (Vue)](https://tanstack.com/table/v8)
*   **Componentes Adicionales**:
    *   **Selector de Fechas**: [V-Calendar](https://vcalendar.io/)
*   **Enrutamiento**: [Vue Router](https://router.vuejs.org/)
*   **Herramientas de Desarrollo**: [Vite](https://vitejs.dev/)


## 🏁 Cómo Empezar

Sigue estos pasos para configurar y ejecutar el proyecto en tu entorno local.

### Prerrequisitos

*   [Node.js](https://nodejs.org/) (versión 18.x o superior recomendada)
*   [npm](https://www.npmjs.com/) o [yarn](https://yarnpkg.com/)

### 🛑 Importante: Configuración del Backend

Este proyecto es únicamente el **frontend**. Requiere un servidor de backend que se ejecute en `http://localhost:8080` y que proporcione un endpoint `/stocks`.

La lógica de fetching se encuentra en `src/stores/stockStore.ts` y espera recibir un array de objetos `StockAction`. Asegúrate de tener tu API de backend en funcionamiento antes de iniciar el frontend.

### Instalación del Frontend

1.  **Clona el repositorio:**
    ```bash
    git clone https://github.com/eduardovparga/stock-app.git
    cd equity-pulse
    ```

2.  **Instala las dependencias:**
    ```bash
    npm install
    ```

3.  **Ejecuta el servidor de desarrollo:**
    ```bash
    npm run dev
    ```

4.  Abre tu navegador y visita `http://localhost:5173` (o el puerto que Vite indique en la consola).

## 📜 Scripts Disponibles

En el `package.json`, encontrarás los siguientes scripts:

*   `npm run dev`: Inicia el servidor de desarrollo de Vite con Hot-Module Replacement (HMR).
*   `npm run build`: Compila y transpila la aplicación para producción en el directorio `dist/`.
*   `npm run preview`: Sirve localmente el build de producción desde el directorio `dist/` para previsualizarlo.
