// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Sidebar() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"hs-application-sidebar\" class=\"hs-overlay  [--auto-close:lg]\n\t\t\t\ths-overlay-open:translate-x-0\n\t\t\t\t-translate-x-full transition-all duration-300 transform\n\t\t\t\tw-[260px] h-full\n\t\t\t\thidden\n\t\t\t\tfixed inset-y-0 start-0 z-[60]\n\t\t\t\tbg-white border-e border-gray-200\n\t\t\t\tlg:block lg:translate-x-0 lg:end-auto lg:bottom-0\n\t\t\t\tdark:bg-neutral-800 dark:border-neutral-700\" role=\"dialog\" tabindex=\"-1\" aria-label=\"Sidebar\"><div class=\"relative flex flex-col h-full max-h-full\"><div class=\"px-6 pt-4\"><!-- Logo --><a class=\"flex-none rounded-xl text-xl inline-block font-semibold focus:outline-none focus:opacity-80\" href=\"#\" aria-label=\"Preline\"><svg width=\"220px\" height=\"100px\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"-65.46668701171876 19.929394531250004 630.9333740234375 110.14121093749999\" style=\"background: rgb(255, 255, 255);\" preserveAspectRatio=\"xMidYMid\"><defs><linearGradient id=\"editing-shadow-gradient1\" x1=\"0\" x2=\"0\" y1=\"0\" y2=\"1\"><stop offset=\"0\" stop-color=\"#534b47\" stop-opacity=\"0\"></stop><stop offset=\"0.2\" stop-color=\"#534b47\" stop-opacity=\"0\"></stop><stop offset=\"1\" stop-color=\"#534b47\" stop-opacity=\"1\"></stop></linearGradient><linearGradient id=\"editing-shadow-gradient2\" x1=\"-0.2771459614569709\" x2=\"1.277145961456971\" y1=\"-0.1293203910498374\" y2=\"1.1293203910498373\"><stop offset=\"0\" stop-color=\"#e2fff7\"></stop><stop offset=\"1\" stop-color=\"#004e72\"></stop></linearGradient><filter id=\"editing-shadow\" x=\"-100%\" y=\"-100%\" width=\"300%\" height=\"300%\"><feGaussianBlur stdDeviation=\"0\"></feGaussianBlur><feComposite operator=\"in\" in2=\"SourceGraphic\"></feComposite><feGaussianBlur stdDeviation=\"0\"></feGaussianBlur></filter></defs><g transform=\"translate(250,75)\"><g transform=\"translate(-21, -4.1)\"><g transform=\"skewX(45)\"><g filter=\"url(#editing-shadow)\"><g transform=\"translate(-191.46500968933105, 24.750001907348633)\"><path d=\"M40.85 0L40.85 0L20.06 0L20.06 0Q11.81 0 7.82-3.30L7.82-3.30L7.82-3.30Q3.83-6.60 3.83-13.46L3.83-13.46L3.83-34.06L3.83-34.06Q3.83-40.92 7.82-44.22L7.82-44.22L7.82-44.22Q11.81-47.52 20.06-47.52L20.06-47.52L36.56-47.52L36.56-47.52Q37.82-47.52 38.38-46.96L38.38-46.96L38.38-46.96Q38.94-46.40 38.94-45.14L38.94-45.14L38.94-37.16L38.94-37.16Q38.94-35.90 38.38-35.34L38.38-35.34L38.38-35.34Q37.82-34.78 36.56-34.78L36.56-34.78L22.84-34.78L22.84-34.78Q20.79-34.78 19.80-33.96L19.80-33.96L19.80-33.96Q18.81-33.13 18.81-31.48L18.81-31.48L18.81-15.71L18.81-15.71Q18.81-14.06 19.57-13.27L19.57-13.27L19.57-13.27Q20.33-12.47 21.85-12.47L21.85-12.47L28.25-12.47L28.25-22.77L28.25-22.77Q28.25-24.02 28.81-24.59L28.81-24.59L28.81-24.59Q29.37-25.15 30.62-25.15L30.62-25.15L40.85-25.15L40.85-25.15Q42.11-25.15 42.67-24.59L42.67-24.59L42.67-24.59Q43.23-24.02 43.23-22.77L43.23-22.77L43.23-2.38L43.23-2.38Q43.23-1.12 42.67-0.56L42.67-0.56L42.67-0.56Q42.11 0 40.85 0ZM22.90-5.94L22.90-5.94L36.04-5.94L36.04-19.21L35.38-19.21L35.38-6.60L22.90-6.60L22.90-6.60Q16.24-6.60 13.96-8.38L13.96-8.38L13.96-8.38Q11.68-10.16 11.68-14.45L11.68-14.45L11.68-32.93L11.68-32.93Q11.68-37.29 13.73-39.11L13.73-39.11L13.73-39.11Q15.77-40.92 21.58-40.92L21.58-40.92L33-40.92L33-41.58L21.58-41.58L21.58-41.58Q15.25-41.58 13.13-39.60L13.13-39.60L13.13-39.60Q11.02-37.62 11.02-32.93L11.02-32.93L11.02-14.45L11.02-14.45Q11.02-12.14 11.52-10.53L11.52-10.53L11.52-10.53Q12.01-8.91 13.33-7.89L13.33-7.89L13.33-7.89Q14.65-6.86 16.96-6.40L16.96-6.40L16.96-6.40Q19.27-5.94 22.90-5.94ZM91.54-34.98L91.54-12.47L91.54-12.47Q91.54-9.44 90.65-6.96L90.65-6.96L90.65-6.96Q89.76-4.49 87.48-2.74L87.48-2.74L87.48-2.74Q85.21-0.99 81.21 0L81.21 0L81.21 0Q77.22 0.99 71.08 0.99L71.08 0.99L71.08 0.99Q64.94 0.99 60.92 0L60.92 0L60.92 0Q56.89-0.99 54.62-2.74L54.62-2.74L54.62-2.74Q52.34-4.49 51.45-6.96L51.45-6.96L51.45-6.96Q50.56-9.44 50.56-12.47L50.56-12.47L50.56-34.98L50.56-34.98Q50.56-38.02 51.45-40.49L51.45-40.49L51.45-40.49Q52.34-42.97 54.62-44.75L54.62-44.75L54.62-44.75Q56.89-46.53 60.92-47.52L60.92-47.52L60.92-47.52Q64.94-48.51 71.08-48.51L71.08-48.51L71.08-48.51Q77.22-48.51 81.21-47.52L81.21-47.52L81.21-47.52Q85.21-46.53 87.48-44.75L87.48-44.75L87.48-44.75Q89.76-42.97 90.65-40.49L90.65-40.49L90.65-40.49Q91.54-38.02 91.54-34.98L91.54-34.98ZM84.35-13.99L84.35-13.99L84.35-33.46L84.35-33.46Q84.35-37.42 81.51-39.90L81.51-39.90L81.51-39.90Q78.67-42.37 70.95-42.37L70.95-42.37L70.95-42.37Q63.23-42.37 60.46-39.90L60.46-39.90L60.46-39.90Q57.68-37.42 57.68-33.46L57.68-33.46L57.68-13.99L57.68-13.99Q57.68-10.03 60.46-7.56L60.46-7.56L60.46-7.56Q63.23-5.08 70.95-5.08L70.95-5.08L70.95-5.08Q78.67-5.08 81.51-7.56L81.51-7.56L81.51-7.56Q84.35-10.03 84.35-13.99ZM58.34-13.99L58.34-13.99L58.34-33.46L58.34-33.46Q58.34-37.16 60.98-39.44L60.98-39.44L60.98-39.44Q63.62-41.71 70.95-41.71L70.95-41.71L70.95-41.71Q78.28-41.71 80.98-39.44L80.98-39.44L80.98-39.44Q83.69-37.16 83.69-33.46L83.69-33.46L83.69-13.99L83.69-13.99Q83.69-10.30 80.98-8.02L80.98-8.02L80.98-8.02Q78.28-5.74 70.95-5.74L70.95-5.74L70.95-5.74Q63.62-5.74 60.98-8.02L60.98-8.02L60.98-8.02Q58.34-10.30 58.34-13.99ZM65.54-32.60L65.54-14.85L65.54-14.85Q65.54-13.53 66.53-12.64L66.53-12.64L66.53-12.64Q67.52-11.75 71.08-11.75L71.08-11.75L71.08-11.75Q74.65-11.75 75.64-12.64L75.64-12.64L75.64-12.64Q76.63-13.53 76.63-14.85L76.63-14.85L76.63-32.60L76.63-32.60Q76.63-33.86 75.64-34.75L75.64-34.75L75.64-34.75Q74.65-35.64 71.08-35.64L71.08-35.64L71.08-35.64Q67.52-35.64 66.53-34.75L66.53-34.75L66.53-34.75Q65.54-33.86 65.54-32.60L65.54-32.60ZM135.89 0L135.89 0L102.30 0L102.30 0Q101.05 0 100.48-0.56L100.48-0.56L100.48-0.56Q99.92-1.12 99.92-2.38L99.92-2.38L99.92-45.14L99.92-45.14Q99.92-46.40 100.48-46.96L100.48-46.96L100.48-46.96Q101.05-47.52 102.30-47.52L102.30-47.52L112.53-47.52L112.53-47.52Q113.78-47.52 114.34-46.96L114.34-46.96L114.34-46.96Q114.91-46.40 114.91-45.14L114.91-45.14L114.91-12.47L123.95-12.47L123.95-23.17L123.95-23.17Q123.95-24.42 124.51-24.98L124.51-24.98L124.51-24.98Q125.07-25.54 126.32-25.54L126.32-25.54L135.89-25.54L135.89-25.54Q137.15-25.54 137.71-24.98L137.71-24.98L137.71-24.98Q138.27-24.42 138.27-23.17L138.27-23.17L138.27-2.38L138.27-2.38Q138.27-1.12 137.71-0.56L137.71-0.56L137.71-0.56Q137.15 0 135.89 0ZM107.12-41.58L107.12-5.94L131.41-5.94L131.41-19.60L130.75-19.60L130.75-6.60L107.78-6.60L107.78-41.58L107.12-41.58ZM152.46-17.95L152.46-20.00L152.46-20.00Q152.46-21.78 152.76-23.30L152.76-23.30L152.76-23.30Q153.05-24.82 153.52-26.14L153.52-26.14L157.74-36.96L157.74-36.96Q158.40-38.61 159.19-39.77L159.19-39.77L159.19-39.77Q159.98-40.92 161.96-40.92L161.96-40.92L168.43-40.92L168.43-40.92Q170.28-40.92 171.07-39.77L171.07-39.77L171.07-39.77Q171.86-38.61 172.59-36.89L172.59-36.89L177.14-26.14L177.14-26.14Q177.61-24.82 177.90-23.30L177.90-23.30L177.90-23.30Q178.20-21.78 178.20-20.00L178.20-20.00L178.20-17.95L152.46-17.95ZM163.02-33.66L163.02-33.66L159.92-23.63L170.28-23.63L167.18-33.66L167.18-33.66Q166.91-34.25 166.62-34.52L166.62-34.52L166.62-34.52Q166.32-34.78 165.92-34.78L165.92-34.78L164.27-34.78L164.27-34.78Q163.88-34.78 163.58-34.52L163.58-34.52L163.58-34.52Q163.28-34.25 163.02-33.66ZM178.20-17.29L178.20-5.94L178.86-5.94L178.86-20.00L178.86-20.00Q178.86-21.85 178.56-23.40L178.56-23.40L178.56-23.40Q178.27-24.95 177.74-26.40L177.74-26.40L173.18-37.16L173.18-37.16Q172.39-39.07 171.50-40.33L171.50-40.33L171.50-40.33Q170.61-41.58 168.43-41.58L168.43-41.58L161.96-41.58L161.96-41.58Q159.72-41.58 158.80-40.33L158.80-40.33L158.80-40.33Q157.87-39.07 157.15-37.22L157.15-37.22L152.92-26.40L152.92-26.40Q152.39-24.95 152.10-23.40L152.10-23.40L152.10-23.40Q151.80-21.85 151.80-20.00L151.80-20.00L151.80-5.94L152.46-5.94L152.46-17.29L178.20-17.29ZM159.26-11.62L159.26-2.38L159.26-2.38Q159.26-1.12 158.70-0.56L158.70-0.56L158.70-0.56Q158.14 0 156.88 0L156.88 0L147.18 0L147.18 0Q145.93 0 145.37-0.56L145.37-0.56L145.37-0.56Q144.80-1.12 144.80-2.38L144.80-2.38L144.80-19.80L144.80-19.80Q144.80-21.78 145.43-24.35L145.43-24.35L145.43-24.35Q146.06-26.93 147.38-30.36L147.38-30.36L153.12-45.34L153.12-45.34Q153.52-46.46 154.34-46.99L154.34-46.99L154.34-46.99Q155.17-47.52 156.49-47.52L156.49-47.52L174.24-47.52L174.24-47.52Q175.49-47.52 176.32-46.99L176.32-46.99L176.32-46.99Q177.14-46.46 177.54-45.34L177.54-45.34L183.28-30.36L183.28-30.36Q184.60-26.93 185.23-24.35L185.23-24.35L185.23-24.35Q185.86-21.78 185.86-19.80L185.86-19.80L185.86-2.38L185.86-2.38Q185.86-1.12 185.30-0.56L185.30-0.56L185.30-0.56Q184.73 0 183.48 0L183.48 0L173.51 0L173.51 0Q172.19 0 171.57-0.56L171.57-0.56L171.57-0.56Q170.94-1.12 170.94-2.38L170.94-2.38L170.94-11.62L159.26-11.62ZM226.97-13.73L202.69-41.05L202.69-41.05Q202.36-41.38 202.16-41.51L202.16-41.51L202.16-41.51Q201.96-41.65 201.70-41.65L201.70-41.65L201.70-41.65Q201.43-41.65 201.17-41.45L201.17-41.45L201.17-41.45Q200.90-41.25 200.90-40.85L200.90-40.85L200.90-5.94L201.56-5.94L201.56-40.72L201.56-40.72Q201.56-40.79 201.60-40.85L201.60-40.85L201.60-40.85Q201.63-40.92 201.70-40.92L201.70-40.92L201.70-40.92Q201.76-40.92 201.86-40.85L201.86-40.85L201.86-40.85Q201.96-40.79 202.03-40.72L202.03-40.72L226.97-12.74L226.97-5.94L227.63-5.94L227.63-41.58L226.97-41.58L226.97-13.73ZM206.25 0L206.25 0L196.35 0L196.35 0Q195.10 0 194.54-0.56L194.54-0.56L194.54-0.56Q193.97-1.12 193.97-2.38L193.97-2.38L193.97-45.14L193.97-45.14Q193.97-46.40 194.54-46.96L194.54-46.96L194.54-46.96Q195.10-47.52 196.35-47.52L196.35-47.52L203.21-47.52L203.21-47.52Q204.47-47.52 205.36-47.12L205.36-47.12L205.36-47.12Q206.25-46.73 207.11-45.74L207.11-45.74L219.91-31.09L219.91-45.14L219.91-45.14Q219.91-46.40 220.47-46.96L220.47-46.96L220.47-46.96Q221.03-47.52 222.29-47.52L222.29-47.52L232.19-47.52L232.19-47.52Q233.44-47.52 234.00-46.96L234.00-46.96L234.00-46.96Q234.56-46.40 234.56-45.14L234.56-45.14L234.56-2.38L234.56-2.38Q234.56-1.12 234.00-0.56L234.00-0.56L234.00-0.56Q233.44 0 232.19 0L232.19 0L222.29 0L222.29 0Q221.03 0 220.47-0.56L220.47-0.56L220.47-0.56Q219.91-1.12 219.91-2.38L219.91-2.38L219.91-10.49L208.63-24.16L208.63-2.38L208.63-2.38Q208.63-1.12 208.07-0.56L208.07-0.56L208.07-0.56Q207.50 0 206.25 0ZM279.97 0L279.97 0L259.18 0L259.18 0Q250.93 0 246.94-3.30L246.94-3.30L246.94-3.30Q242.95-6.60 242.95-13.46L242.95-13.46L242.95-34.06L242.95-34.06Q242.95-40.92 246.94-44.22L246.94-44.22L246.94-44.22Q250.93-47.52 259.18-47.52L259.18-47.52L275.68-47.52L275.68-47.52Q276.94-47.52 277.50-46.96L277.50-46.96L277.50-46.96Q278.06-46.40 278.06-45.14L278.06-45.14L278.06-37.16L278.06-37.16Q278.06-35.90 277.50-35.34L277.50-35.34L277.50-35.34Q276.94-34.78 275.68-34.78L275.68-34.78L261.95-34.78L261.95-34.78Q259.91-34.78 258.92-33.96L258.92-33.96L258.92-33.96Q257.93-33.13 257.93-31.48L257.93-31.48L257.93-15.71L257.93-15.71Q257.93-14.06 258.69-13.27L258.69-13.27L258.69-13.27Q259.45-12.47 260.96-12.47L260.96-12.47L267.37-12.47L267.37-22.77L267.37-22.77Q267.37-24.02 267.93-24.59L267.93-24.59L267.93-24.59Q268.49-25.15 269.74-25.15L269.74-25.15L279.97-25.15L279.97-25.15Q281.23-25.15 281.79-24.59L281.79-24.59L281.79-24.59Q282.35-24.02 282.35-22.77L282.35-22.77L282.35-2.38L282.35-2.38Q282.35-1.12 281.79-0.56L281.79-0.56L281.79-0.56Q281.23 0 279.97 0ZM262.02-5.94L262.02-5.94L275.15-5.94L275.15-19.21L274.49-19.21L274.49-6.60L262.02-6.60L262.02-6.60Q255.35-6.60 253.08-8.38L253.08-8.38L253.08-8.38Q250.80-10.16 250.80-14.45L250.80-14.45L250.80-32.93L250.80-32.93Q250.80-37.29 252.85-39.11L252.85-39.11L252.85-39.11Q254.89-40.92 260.70-40.92L260.70-40.92L272.12-40.92L272.12-41.58L260.70-41.58L260.70-41.58Q254.36-41.58 252.25-39.60L252.25-39.60L252.25-39.60Q250.14-37.62 250.14-32.93L250.14-32.93L250.14-14.45L250.14-14.45Q250.14-12.14 250.64-10.53L250.64-10.53L250.64-10.53Q251.13-8.91 252.45-7.89L252.45-7.89L252.45-7.89Q253.77-6.86 256.08-6.40L256.08-6.40L256.08-6.40Q258.39-5.94 262.02-5.94ZM318.12 0L318.12 0L292.78 0L292.78 0Q291.52 0 290.96-0.56L290.96-0.56L290.96-0.56Q290.40-1.12 290.40-2.38L290.40-2.38L290.40-45.14L290.40-45.14Q290.40-46.40 290.96-46.96L290.96-46.96L290.96-46.96Q291.52-47.52 292.78-47.52L292.78-47.52L313.83-47.52L313.83-47.52Q320.30-47.52 323.40-45.05L323.40-45.05L323.40-45.05Q326.50-42.57 326.50-37.42L326.50-37.42L326.50-34.25L326.50-34.25Q326.50-31.48 325.55-29.73L325.55-29.73L325.55-29.73Q324.59-27.98 322.61-27.13L322.61-27.13L322.61-27.13Q326.83-26.73 328.88-24.22L328.88-24.22L328.88-24.22Q330.92-21.71 330.92-16.90L330.92-16.90L330.92-10.82L330.92-10.82Q330.92-5.28 327.79-2.64L327.79-2.64L327.79-2.64Q324.65 0 318.12 0ZM297.73-41.58L297.73-5.94L315.48-5.94L315.48-5.94Q318.25-5.94 319.97-6.44L319.97-6.44L319.97-6.44Q321.68-6.93 322.57-7.79L322.57-7.79L322.57-7.79Q323.47-8.65 323.80-9.87L323.80-9.87L323.80-9.87Q324.13-11.09 324.13-12.54L324.13-12.54L324.13-16.76L324.13-16.76Q324.13-18.41 323.83-19.80L323.83-19.80L323.83-19.80Q323.53-21.19 322.61-22.21L322.61-22.21L322.61-22.21Q321.68-23.23 319.97-23.79L319.97-23.79L319.97-23.79Q318.25-24.35 315.55-24.35L315.55-24.35L315.55-24.35Q318.38-25.08 319.08-26.73L319.08-26.73L319.08-26.73Q319.77-28.38 319.77-30.36L319.77-30.36L319.77-34.58L319.77-34.58Q319.77-36.04 319.44-37.29L319.44-37.29L319.44-37.29Q319.11-38.54 318.19-39.50L318.19-39.50L318.19-39.50Q317.26-40.46 315.55-41.02L315.55-41.02L315.55-41.02Q313.83-41.58 311.12-41.58L311.12-41.58L297.73-41.58ZM298.39-40.92L311.12-40.92L311.12-40.92Q316.07-40.92 317.59-39.07L317.59-39.07L317.59-39.07Q319.11-37.22 319.11-34.58L319.11-34.58L319.11-30.36L319.11-30.36Q319.11-29.17 318.88-28.08L318.88-28.08L318.88-28.08Q318.65-26.99 317.82-26.17L317.82-26.17L317.82-26.17Q317.00-25.34 315.41-24.85L315.41-24.85L315.41-24.85Q313.83-24.35 311.12-24.35L311.12-24.35L298.39-24.35L298.39-40.92ZM310.86-36.30L305.25-36.30L305.25-29.17L310.86-29.17L310.86-29.17Q312.11-29.17 312.74-29.83L312.74-29.83L312.74-29.83Q313.37-30.49 313.37-31.81L313.37-31.81L313.37-33.66L313.37-33.66Q313.37-34.98 312.74-35.64L312.74-35.64L312.74-35.64Q312.11-36.30 310.86-36.30L310.86-36.30ZM298.39-23.69L315.55-23.69L315.55-23.69Q320.50-23.69 321.98-21.71L321.98-21.71L321.98-21.71Q323.47-19.73 323.47-16.76L323.47-16.76L323.47-12.54L323.47-12.54Q323.47-9.83 321.95-8.22L321.95-8.22L321.95-8.22Q320.43-6.60 315.48-6.60L315.48-6.60L298.39-6.60L298.39-23.69ZM313.83-19.01L305.25-19.01L305.25-11.22L313.83-11.22L313.83-11.22Q315.08-11.22 315.71-11.88L315.71-11.88L315.71-11.88Q316.34-12.54 316.34-13.86L316.34-13.86L316.34-16.43L316.34-16.43Q316.34-17.75 315.71-18.38L315.71-18.38L315.71-18.38Q315.08-19.01 313.83-19.01L313.83-19.01ZM362.87 0L362.87 0L340.63 0L340.63 0Q339.37 0 338.81-0.56L338.81-0.56L338.81-0.56Q338.25-1.12 338.25-2.38L338.25-2.38L338.25-45.14L338.25-45.14Q338.25-46.40 338.81-46.96L338.81-46.96L338.81-46.96Q339.37-47.52 340.63-47.52L340.63-47.52L362.87-47.52L362.87-47.52Q371.12-47.52 375.11-44.22L375.11-44.22L375.11-44.22Q379.10-40.92 379.10-34.06L379.10-34.06L379.10-13.46L379.10-13.46Q379.10-6.60 375.11-3.30L375.11-3.30L375.11-3.30Q371.12 0 362.87 0ZM345.44-41.58L345.44-5.94L361.42-5.94L361.42-5.94Q364.78-5.94 366.89-6.63L366.89-6.63L366.89-6.63Q369.01-7.33 370.13-8.51L370.13-8.51L370.13-8.51Q371.25-9.70 371.61-11.35L371.61-11.35L371.61-11.35Q371.98-13.00 371.98-14.92L371.98-14.92L371.98-32.41L371.98-32.41Q371.98-34.32 371.58-36.00L371.58-36.00L371.58-36.00Q371.18-37.69 370.06-38.91L370.06-38.91L370.06-38.91Q368.94-40.13 366.86-40.85L366.86-40.85L366.86-40.85Q364.78-41.58 361.42-41.58L361.42-41.58L345.44-41.58ZM346.10-40.92L361.42-40.92L361.42-40.92Q364.52-40.92 366.47-40.26L366.47-40.26L366.47-40.26Q368.41-39.60 369.47-38.48L369.47-38.48L369.47-38.48Q370.52-37.36 370.92-35.77L370.92-35.77L370.92-35.77Q371.32-34.19 371.32-32.41L371.32-32.41L371.32-14.92L371.32-14.92Q371.32-13.13 370.95-11.58L370.95-11.58L370.95-11.58Q370.59-10.03 369.53-8.94L369.53-8.94L369.53-8.94Q368.48-7.85 366.53-7.23L366.53-7.23L366.53-7.23Q364.58-6.60 361.42-6.60L361.42-6.60L346.10-6.60L346.10-40.92ZM361.09-34.91L353.23-34.91L353.23-12.47L361.09-12.47L361.09-12.47Q362.67-12.47 363.43-13.27L363.43-13.27L363.43-13.27Q364.19-14.06 364.19-15.71L364.19-15.71L364.19-31.75L364.19-31.75Q364.19-33.40 363.43-34.16L363.43-34.16L363.43-34.16Q362.67-34.91 361.09-34.91L361.09-34.91Z\" fill=\"url(#editing-shadow-gradient1)\"></path></g></g></g></g><g><g transform=\"translate(-191.46500968933105, 24.750001907348633)\"><path d=\"M40.85 0L40.85 0L20.06 0L20.06 0Q11.81 0 7.82-3.30L7.82-3.30L7.82-3.30Q3.83-6.60 3.83-13.46L3.83-13.46L3.83-34.06L3.83-34.06Q3.83-40.92 7.82-44.22L7.82-44.22L7.82-44.22Q11.81-47.52 20.06-47.52L20.06-47.52L36.56-47.52L36.56-47.52Q37.82-47.52 38.38-46.96L38.38-46.96L38.38-46.96Q38.94-46.40 38.94-45.14L38.94-45.14L38.94-37.16L38.94-37.16Q38.94-35.90 38.38-35.34L38.38-35.34L38.38-35.34Q37.82-34.78 36.56-34.78L36.56-34.78L22.84-34.78L22.84-34.78Q20.79-34.78 19.80-33.96L19.80-33.96L19.80-33.96Q18.81-33.13 18.81-31.48L18.81-31.48L18.81-15.71L18.81-15.71Q18.81-14.06 19.57-13.27L19.57-13.27L19.57-13.27Q20.33-12.47 21.85-12.47L21.85-12.47L28.25-12.47L28.25-22.77L28.25-22.77Q28.25-24.02 28.81-24.59L28.81-24.59L28.81-24.59Q29.37-25.15 30.62-25.15L30.62-25.15L40.85-25.15L40.85-25.15Q42.11-25.15 42.67-24.59L42.67-24.59L42.67-24.59Q43.23-24.02 43.23-22.77L43.23-22.77L43.23-2.38L43.23-2.38Q43.23-1.12 42.67-0.56L42.67-0.56L42.67-0.56Q42.11 0 40.85 0ZM22.90-5.94L22.90-5.94L36.04-5.94L36.04-19.21L35.38-19.21L35.38-6.60L22.90-6.60L22.90-6.60Q16.24-6.60 13.96-8.38L13.96-8.38L13.96-8.38Q11.68-10.16 11.68-14.45L11.68-14.45L11.68-32.93L11.68-32.93Q11.68-37.29 13.73-39.11L13.73-39.11L13.73-39.11Q15.77-40.92 21.58-40.92L21.58-40.92L33-40.92L33-41.58L21.58-41.58L21.58-41.58Q15.25-41.58 13.13-39.60L13.13-39.60L13.13-39.60Q11.02-37.62 11.02-32.93L11.02-32.93L11.02-14.45L11.02-14.45Q11.02-12.14 11.52-10.53L11.52-10.53L11.52-10.53Q12.01-8.91 13.33-7.89L13.33-7.89L13.33-7.89Q14.65-6.86 16.96-6.40L16.96-6.40L16.96-6.40Q19.27-5.94 22.90-5.94ZM91.54-34.98L91.54-12.47L91.54-12.47Q91.54-9.44 90.65-6.96L90.65-6.96L90.65-6.96Q89.76-4.49 87.48-2.74L87.48-2.74L87.48-2.74Q85.21-0.99 81.21 0L81.21 0L81.21 0Q77.22 0.99 71.08 0.99L71.08 0.99L71.08 0.99Q64.94 0.99 60.92 0L60.92 0L60.92 0Q56.89-0.99 54.62-2.74L54.62-2.74L54.62-2.74Q52.34-4.49 51.45-6.96L51.45-6.96L51.45-6.96Q50.56-9.44 50.56-12.47L50.56-12.47L50.56-34.98L50.56-34.98Q50.56-38.02 51.45-40.49L51.45-40.49L51.45-40.49Q52.34-42.97 54.62-44.75L54.62-44.75L54.62-44.75Q56.89-46.53 60.92-47.52L60.92-47.52L60.92-47.52Q64.94-48.51 71.08-48.51L71.08-48.51L71.08-48.51Q77.22-48.51 81.21-47.52L81.21-47.52L81.21-47.52Q85.21-46.53 87.48-44.75L87.48-44.75L87.48-44.75Q89.76-42.97 90.65-40.49L90.65-40.49L90.65-40.49Q91.54-38.02 91.54-34.98L91.54-34.98ZM84.35-13.99L84.35-13.99L84.35-33.46L84.35-33.46Q84.35-37.42 81.51-39.90L81.51-39.90L81.51-39.90Q78.67-42.37 70.95-42.37L70.95-42.37L70.95-42.37Q63.23-42.37 60.46-39.90L60.46-39.90L60.46-39.90Q57.68-37.42 57.68-33.46L57.68-33.46L57.68-13.99L57.68-13.99Q57.68-10.03 60.46-7.56L60.46-7.56L60.46-7.56Q63.23-5.08 70.95-5.08L70.95-5.08L70.95-5.08Q78.67-5.08 81.51-7.56L81.51-7.56L81.51-7.56Q84.35-10.03 84.35-13.99ZM58.34-13.99L58.34-13.99L58.34-33.46L58.34-33.46Q58.34-37.16 60.98-39.44L60.98-39.44L60.98-39.44Q63.62-41.71 70.95-41.71L70.95-41.71L70.95-41.71Q78.28-41.71 80.98-39.44L80.98-39.44L80.98-39.44Q83.69-37.16 83.69-33.46L83.69-33.46L83.69-13.99L83.69-13.99Q83.69-10.30 80.98-8.02L80.98-8.02L80.98-8.02Q78.28-5.74 70.95-5.74L70.95-5.74L70.95-5.74Q63.62-5.74 60.98-8.02L60.98-8.02L60.98-8.02Q58.34-10.30 58.34-13.99ZM65.54-32.60L65.54-14.85L65.54-14.85Q65.54-13.53 66.53-12.64L66.53-12.64L66.53-12.64Q67.52-11.75 71.08-11.75L71.08-11.75L71.08-11.75Q74.65-11.75 75.64-12.64L75.64-12.64L75.64-12.64Q76.63-13.53 76.63-14.85L76.63-14.85L76.63-32.60L76.63-32.60Q76.63-33.86 75.64-34.75L75.64-34.75L75.64-34.75Q74.65-35.64 71.08-35.64L71.08-35.64L71.08-35.64Q67.52-35.64 66.53-34.75L66.53-34.75L66.53-34.75Q65.54-33.86 65.54-32.60L65.54-32.60ZM135.89 0L135.89 0L102.30 0L102.30 0Q101.05 0 100.48-0.56L100.48-0.56L100.48-0.56Q99.92-1.12 99.92-2.38L99.92-2.38L99.92-45.14L99.92-45.14Q99.92-46.40 100.48-46.96L100.48-46.96L100.48-46.96Q101.05-47.52 102.30-47.52L102.30-47.52L112.53-47.52L112.53-47.52Q113.78-47.52 114.34-46.96L114.34-46.96L114.34-46.96Q114.91-46.40 114.91-45.14L114.91-45.14L114.91-12.47L123.95-12.47L123.95-23.17L123.95-23.17Q123.95-24.42 124.51-24.98L124.51-24.98L124.51-24.98Q125.07-25.54 126.32-25.54L126.32-25.54L135.89-25.54L135.89-25.54Q137.15-25.54 137.71-24.98L137.71-24.98L137.71-24.98Q138.27-24.42 138.27-23.17L138.27-23.17L138.27-2.38L138.27-2.38Q138.27-1.12 137.71-0.56L137.71-0.56L137.71-0.56Q137.15 0 135.89 0ZM107.12-41.58L107.12-5.94L131.41-5.94L131.41-19.60L130.75-19.60L130.75-6.60L107.78-6.60L107.78-41.58L107.12-41.58ZM152.46-17.95L152.46-20.00L152.46-20.00Q152.46-21.78 152.76-23.30L152.76-23.30L152.76-23.30Q153.05-24.82 153.52-26.14L153.52-26.14L157.74-36.96L157.74-36.96Q158.40-38.61 159.19-39.77L159.19-39.77L159.19-39.77Q159.98-40.92 161.96-40.92L161.96-40.92L168.43-40.92L168.43-40.92Q170.28-40.92 171.07-39.77L171.07-39.77L171.07-39.77Q171.86-38.61 172.59-36.89L172.59-36.89L177.14-26.14L177.14-26.14Q177.61-24.82 177.90-23.30L177.90-23.30L177.90-23.30Q178.20-21.78 178.20-20.00L178.20-20.00L178.20-17.95L152.46-17.95ZM163.02-33.66L163.02-33.66L159.92-23.63L170.28-23.63L167.18-33.66L167.18-33.66Q166.91-34.25 166.62-34.52L166.62-34.52L166.62-34.52Q166.32-34.78 165.92-34.78L165.92-34.78L164.27-34.78L164.27-34.78Q163.88-34.78 163.58-34.52L163.58-34.52L163.58-34.52Q163.28-34.25 163.02-33.66ZM178.20-17.29L178.20-5.94L178.86-5.94L178.86-20.00L178.86-20.00Q178.86-21.85 178.56-23.40L178.56-23.40L178.56-23.40Q178.27-24.95 177.74-26.40L177.74-26.40L173.18-37.16L173.18-37.16Q172.39-39.07 171.50-40.33L171.50-40.33L171.50-40.33Q170.61-41.58 168.43-41.58L168.43-41.58L161.96-41.58L161.96-41.58Q159.72-41.58 158.80-40.33L158.80-40.33L158.80-40.33Q157.87-39.07 157.15-37.22L157.15-37.22L152.92-26.40L152.92-26.40Q152.39-24.95 152.10-23.40L152.10-23.40L152.10-23.40Q151.80-21.85 151.80-20.00L151.80-20.00L151.80-5.94L152.46-5.94L152.46-17.29L178.20-17.29ZM159.26-11.62L159.26-2.38L159.26-2.38Q159.26-1.12 158.70-0.56L158.70-0.56L158.70-0.56Q158.14 0 156.88 0L156.88 0L147.18 0L147.18 0Q145.93 0 145.37-0.56L145.37-0.56L145.37-0.56Q144.80-1.12 144.80-2.38L144.80-2.38L144.80-19.80L144.80-19.80Q144.80-21.78 145.43-24.35L145.43-24.35L145.43-24.35Q146.06-26.93 147.38-30.36L147.38-30.36L153.12-45.34L153.12-45.34Q153.52-46.46 154.34-46.99L154.34-46.99L154.34-46.99Q155.17-47.52 156.49-47.52L156.49-47.52L174.24-47.52L174.24-47.52Q175.49-47.52 176.32-46.99L176.32-46.99L176.32-46.99Q177.14-46.46 177.54-45.34L177.54-45.34L183.28-30.36L183.28-30.36Q184.60-26.93 185.23-24.35L185.23-24.35L185.23-24.35Q185.86-21.78 185.86-19.80L185.86-19.80L185.86-2.38L185.86-2.38Q185.86-1.12 185.30-0.56L185.30-0.56L185.30-0.56Q184.73 0 183.48 0L183.48 0L173.51 0L173.51 0Q172.19 0 171.57-0.56L171.57-0.56L171.57-0.56Q170.94-1.12 170.94-2.38L170.94-2.38L170.94-11.62L159.26-11.62ZM226.97-13.73L202.69-41.05L202.69-41.05Q202.36-41.38 202.16-41.51L202.16-41.51L202.16-41.51Q201.96-41.65 201.70-41.65L201.70-41.65L201.70-41.65Q201.43-41.65 201.17-41.45L201.17-41.45L201.17-41.45Q200.90-41.25 200.90-40.85L200.90-40.85L200.90-5.94L201.56-5.94L201.56-40.72L201.56-40.72Q201.56-40.79 201.60-40.85L201.60-40.85L201.60-40.85Q201.63-40.92 201.70-40.92L201.70-40.92L201.70-40.92Q201.76-40.92 201.86-40.85L201.86-40.85L201.86-40.85Q201.96-40.79 202.03-40.72L202.03-40.72L226.97-12.74L226.97-5.94L227.63-5.94L227.63-41.58L226.97-41.58L226.97-13.73ZM206.25 0L206.25 0L196.35 0L196.35 0Q195.10 0 194.54-0.56L194.54-0.56L194.54-0.56Q193.97-1.12 193.97-2.38L193.97-2.38L193.97-45.14L193.97-45.14Q193.97-46.40 194.54-46.96L194.54-46.96L194.54-46.96Q195.10-47.52 196.35-47.52L196.35-47.52L203.21-47.52L203.21-47.52Q204.47-47.52 205.36-47.12L205.36-47.12L205.36-47.12Q206.25-46.73 207.11-45.74L207.11-45.74L219.91-31.09L219.91-45.14L219.91-45.14Q219.91-46.40 220.47-46.96L220.47-46.96L220.47-46.96Q221.03-47.52 222.29-47.52L222.29-47.52L232.19-47.52L232.19-47.52Q233.44-47.52 234.00-46.96L234.00-46.96L234.00-46.96Q234.56-46.40 234.56-45.14L234.56-45.14L234.56-2.38L234.56-2.38Q234.56-1.12 234.00-0.56L234.00-0.56L234.00-0.56Q233.44 0 232.19 0L232.19 0L222.29 0L222.29 0Q221.03 0 220.47-0.56L220.47-0.56L220.47-0.56Q219.91-1.12 219.91-2.38L219.91-2.38L219.91-10.49L208.63-24.16L208.63-2.38L208.63-2.38Q208.63-1.12 208.07-0.56L208.07-0.56L208.07-0.56Q207.50 0 206.25 0ZM279.97 0L279.97 0L259.18 0L259.18 0Q250.93 0 246.94-3.30L246.94-3.30L246.94-3.30Q242.95-6.60 242.95-13.46L242.95-13.46L242.95-34.06L242.95-34.06Q242.95-40.92 246.94-44.22L246.94-44.22L246.94-44.22Q250.93-47.52 259.18-47.52L259.18-47.52L275.68-47.52L275.68-47.52Q276.94-47.52 277.50-46.96L277.50-46.96L277.50-46.96Q278.06-46.40 278.06-45.14L278.06-45.14L278.06-37.16L278.06-37.16Q278.06-35.90 277.50-35.34L277.50-35.34L277.50-35.34Q276.94-34.78 275.68-34.78L275.68-34.78L261.95-34.78L261.95-34.78Q259.91-34.78 258.92-33.96L258.92-33.96L258.92-33.96Q257.93-33.13 257.93-31.48L257.93-31.48L257.93-15.71L257.93-15.71Q257.93-14.06 258.69-13.27L258.69-13.27L258.69-13.27Q259.45-12.47 260.96-12.47L260.96-12.47L267.37-12.47L267.37-22.77L267.37-22.77Q267.37-24.02 267.93-24.59L267.93-24.59L267.93-24.59Q268.49-25.15 269.74-25.15L269.74-25.15L279.97-25.15L279.97-25.15Q281.23-25.15 281.79-24.59L281.79-24.59L281.79-24.59Q282.35-24.02 282.35-22.77L282.35-22.77L282.35-2.38L282.35-2.38Q282.35-1.12 281.79-0.56L281.79-0.56L281.79-0.56Q281.23 0 279.97 0ZM262.02-5.94L262.02-5.94L275.15-5.94L275.15-19.21L274.49-19.21L274.49-6.60L262.02-6.60L262.02-6.60Q255.35-6.60 253.08-8.38L253.08-8.38L253.08-8.38Q250.80-10.16 250.80-14.45L250.80-14.45L250.80-32.93L250.80-32.93Q250.80-37.29 252.85-39.11L252.85-39.11L252.85-39.11Q254.89-40.92 260.70-40.92L260.70-40.92L272.12-40.92L272.12-41.58L260.70-41.58L260.70-41.58Q254.36-41.58 252.25-39.60L252.25-39.60L252.25-39.60Q250.14-37.62 250.14-32.93L250.14-32.93L250.14-14.45L250.14-14.45Q250.14-12.14 250.64-10.53L250.64-10.53L250.64-10.53Q251.13-8.91 252.45-7.89L252.45-7.89L252.45-7.89Q253.77-6.86 256.08-6.40L256.08-6.40L256.08-6.40Q258.39-5.94 262.02-5.94ZM318.12 0L318.12 0L292.78 0L292.78 0Q291.52 0 290.96-0.56L290.96-0.56L290.96-0.56Q290.40-1.12 290.40-2.38L290.40-2.38L290.40-45.14L290.40-45.14Q290.40-46.40 290.96-46.96L290.96-46.96L290.96-46.96Q291.52-47.52 292.78-47.52L292.78-47.52L313.83-47.52L313.83-47.52Q320.30-47.52 323.40-45.05L323.40-45.05L323.40-45.05Q326.50-42.57 326.50-37.42L326.50-37.42L326.50-34.25L326.50-34.25Q326.50-31.48 325.55-29.73L325.55-29.73L325.55-29.73Q324.59-27.98 322.61-27.13L322.61-27.13L322.61-27.13Q326.83-26.73 328.88-24.22L328.88-24.22L328.88-24.22Q330.92-21.71 330.92-16.90L330.92-16.90L330.92-10.82L330.92-10.82Q330.92-5.28 327.79-2.64L327.79-2.64L327.79-2.64Q324.65 0 318.12 0ZM297.73-41.58L297.73-5.94L315.48-5.94L315.48-5.94Q318.25-5.94 319.97-6.44L319.97-6.44L319.97-6.44Q321.68-6.93 322.57-7.79L322.57-7.79L322.57-7.79Q323.47-8.65 323.80-9.87L323.80-9.87L323.80-9.87Q324.13-11.09 324.13-12.54L324.13-12.54L324.13-16.76L324.13-16.76Q324.13-18.41 323.83-19.80L323.83-19.80L323.83-19.80Q323.53-21.19 322.61-22.21L322.61-22.21L322.61-22.21Q321.68-23.23 319.97-23.79L319.97-23.79L319.97-23.79Q318.25-24.35 315.55-24.35L315.55-24.35L315.55-24.35Q318.38-25.08 319.08-26.73L319.08-26.73L319.08-26.73Q319.77-28.38 319.77-30.36L319.77-30.36L319.77-34.58L319.77-34.58Q319.77-36.04 319.44-37.29L319.44-37.29L319.44-37.29Q319.11-38.54 318.19-39.50L318.19-39.50L318.19-39.50Q317.26-40.46 315.55-41.02L315.55-41.02L315.55-41.02Q313.83-41.58 311.12-41.58L311.12-41.58L297.73-41.58ZM298.39-40.92L311.12-40.92L311.12-40.92Q316.07-40.92 317.59-39.07L317.59-39.07L317.59-39.07Q319.11-37.22 319.11-34.58L319.11-34.58L319.11-30.36L319.11-30.36Q319.11-29.17 318.88-28.08L318.88-28.08L318.88-28.08Q318.65-26.99 317.82-26.17L317.82-26.17L317.82-26.17Q317.00-25.34 315.41-24.85L315.41-24.85L315.41-24.85Q313.83-24.35 311.12-24.35L311.12-24.35L298.39-24.35L298.39-40.92ZM310.86-36.30L305.25-36.30L305.25-29.17L310.86-29.17L310.86-29.17Q312.11-29.17 312.74-29.83L312.74-29.83L312.74-29.83Q313.37-30.49 313.37-31.81L313.37-31.81L313.37-33.66L313.37-33.66Q313.37-34.98 312.74-35.64L312.74-35.64L312.74-35.64Q312.11-36.30 310.86-36.30L310.86-36.30ZM298.39-23.69L315.55-23.69L315.55-23.69Q320.50-23.69 321.98-21.71L321.98-21.71L321.98-21.71Q323.47-19.73 323.47-16.76L323.47-16.76L323.47-12.54L323.47-12.54Q323.47-9.83 321.95-8.22L321.95-8.22L321.95-8.22Q320.43-6.60 315.48-6.60L315.48-6.60L298.39-6.60L298.39-23.69ZM313.83-19.01L305.25-19.01L305.25-11.22L313.83-11.22L313.83-11.22Q315.08-11.22 315.71-11.88L315.71-11.88L315.71-11.88Q316.34-12.54 316.34-13.86L316.34-13.86L316.34-16.43L316.34-16.43Q316.34-17.75 315.71-18.38L315.71-18.38L315.71-18.38Q315.08-19.01 313.83-19.01L313.83-19.01ZM362.87 0L362.87 0L340.63 0L340.63 0Q339.37 0 338.81-0.56L338.81-0.56L338.81-0.56Q338.25-1.12 338.25-2.38L338.25-2.38L338.25-45.14L338.25-45.14Q338.25-46.40 338.81-46.96L338.81-46.96L338.81-46.96Q339.37-47.52 340.63-47.52L340.63-47.52L362.87-47.52L362.87-47.52Q371.12-47.52 375.11-44.22L375.11-44.22L375.11-44.22Q379.10-40.92 379.10-34.06L379.10-34.06L379.10-13.46L379.10-13.46Q379.10-6.60 375.11-3.30L375.11-3.30L375.11-3.30Q371.12 0 362.87 0ZM345.44-41.58L345.44-5.94L361.42-5.94L361.42-5.94Q364.78-5.94 366.89-6.63L366.89-6.63L366.89-6.63Q369.01-7.33 370.13-8.51L370.13-8.51L370.13-8.51Q371.25-9.70 371.61-11.35L371.61-11.35L371.61-11.35Q371.98-13.00 371.98-14.92L371.98-14.92L371.98-32.41L371.98-32.41Q371.98-34.32 371.58-36.00L371.58-36.00L371.58-36.00Q371.18-37.69 370.06-38.91L370.06-38.91L370.06-38.91Q368.94-40.13 366.86-40.85L366.86-40.85L366.86-40.85Q364.78-41.58 361.42-41.58L361.42-41.58L345.44-41.58ZM346.10-40.92L361.42-40.92L361.42-40.92Q364.52-40.92 366.47-40.26L366.47-40.26L366.47-40.26Q368.41-39.60 369.47-38.48L369.47-38.48L369.47-38.48Q370.52-37.36 370.92-35.77L370.92-35.77L370.92-35.77Q371.32-34.19 371.32-32.41L371.32-32.41L371.32-14.92L371.32-14.92Q371.32-13.13 370.95-11.58L370.95-11.58L370.95-11.58Q370.59-10.03 369.53-8.94L369.53-8.94L369.53-8.94Q368.48-7.85 366.53-7.23L366.53-7.23L366.53-7.23Q364.58-6.60 361.42-6.60L361.42-6.60L346.10-6.60L346.10-40.92ZM361.09-34.91L353.23-34.91L353.23-12.47L361.09-12.47L361.09-12.47Q362.67-12.47 363.43-13.27L363.43-13.27L363.43-13.27Q364.19-14.06 364.19-15.71L364.19-15.71L364.19-31.75L364.19-31.75Q364.19-33.40 363.43-34.16L363.43-34.16L363.43-34.16Q362.67-34.91 361.09-34.91L361.09-34.91Z\" fill=\"url(#editing-shadow-gradient2)\"></path></g></g></g><style>text {\n\t\t\t\t\tfont-size: 12px;\n\t\t\t\t\tfont-family: Arial Black;\n\t\t\t\t\tdominant-baseline: central;\n\t\t\t\t\ttext-anchor: middle;\n\t\t\t\t\t}</style></svg></a><!-- End Logo --></div><!-- Content --><div class=\"h-full overflow-y-auto [&amp;::-webkit-scrollbar]:w-2 [&amp;::-webkit-scrollbar-thumb]:rounded-full [&amp;::-webkit-scrollbar-track]:bg-gray-100 [&amp;::-webkit-scrollbar-thumb]:bg-gray-300 dark:[&amp;::-webkit-scrollbar-track]:bg-neutral-700 dark:[&amp;::-webkit-scrollbar-thumb]:bg-neutral-500\"><nav class=\"hs-accordion-group p-3 w-full flex flex-col flex-wrap\" data-hs-accordion-always-open><ul class=\"flex flex-col space-y-1\"><li><a href=\"/dashboard\" class=\"flex items-center gap-x-3.5 py-2 px-2.5 bg-gray-100 text-sm text-gray-800 rounded-lg hover:bg-gray-100 focus:outline-none focus:bg-gray-100 dark:bg-neutral-700 dark:text-white\" href=\"#\"><svg class=\"shrink-0 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z\"></path><polyline points=\"9 22 9 12 15 12 15 22\"></polyline></svg> Dashboard</a></li><li><a href=\"/user\" class=\"flex items-center gap-x-3.5 py-2 px-2.5 bg-gray-100 text-sm text-gray-800 rounded-lg hover:bg-gray-100 focus:outline-none focus:bg-gray-100 dark:bg-neutral-700 dark:text-white\" href=\"#\"><svg class=\"shrink-0 size-4\" xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z\"></path><polyline points=\"9 22 9 12 15 12 15 22\"></polyline></svg> Users</a></li></ul></nav></div><!-- End Content --></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate