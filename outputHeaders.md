# Waifuismo

## Non-Fiction Format for LaTeX [Folder]

**General Non-Fiction (LaTeX)**

  ##  [Text]

About This Template

The goal of this template is make the composition and drafting phase as simple and elegant as possible. It is not meant to turn Scrivener, a rich text editor, into a full-blown LaTeX generator (you'll need to use MultiMarkdown or Pandoc for that; consider using the General Non-Fiction template if that is what you are looking for). Instead, this project contains a number of convenience features that could be applied to your work as you wish in an *à**à** la carte* fashion, to alleviate how much code you'd need to type in yourself. For example, if you find handling images to be cumbersome in LaTeX, you can safely leave that up to the template.

This template caters to those that by and large prefer to compose directly in LaTeX itself, and will be approaching Scrivener primarily as an organisational and composition tool for constructing a longer .tex file from the smaller snippets written into the Draft folder.

When compiled (**File****▸****?****Compile...**), this project will generate a bare bones .tex document(using the flexible Memoir class) suitable for many types of work, fiction and non-fiction alike. The template achieves this by using advanced features available to the Plain Text compile file type, including providing the default extension of ".tex" instead of ".txt".

    ##  [Text]

How To Use This Template

Inside the , create a new folder for each chapter and title each folder with the name of the chapter. (LaTeX itself will of course handle all numbering schemes.) The first chapter folder has been created for you with the placeholder title “Chapter”.

Create a new text document for each section inside the chapter folders. Sections will be numbered and titled by LaTeX using hierarchical numbering (1.1, 1.2, 1.3).

Nest documents beneath documents to create numbered subsections, subsubsections and unnumbered paragraphs, at 3rd, 4th and 5th levels of hierarchy respectively.

You can of course also nest folders if you prefer.

If you need to create the occasional unnumbered section, the Informal Chapter and Informal Section Types can be manually applied to those sections.

If you don’t require a foreword, move the “Foreword” document in the  to the Trash folder. Alternatively, rename it “Preface”, “Abstract” or “Introduction” as you require. Adjust the title, table of contents and other front matter material within this folder, placing any additional front matter sections between the “Begin Document” and “Begin Main Matter” items.

“Notes” and “Ideas” folders have been provided for your convenience, although you can replace them or create different top-level folders for your research materials if necessary (these are just regular folders that have had custom icons assigned to them using the **Documents****▸****?****Change Icon** feature).

    ##  [Text]

General Writing Tips

Foremost, any LaTeX you type in yourself will be dutifully conveyed to the output. As such, you can treat the main editor in Scrivener much like you would a dedicated text editor such as TeXShop. You may still prefer to use another editor to construct LaTeX syntax more easily, or take advantage of the syntax highlighting these types of editors can afford. Simply copy and paste the results of your writings into Scrivener when you’re ready to “save” them.

    ##  [Text]

Adjusting Outline Structure

Scrivener is designed so that you can use the Draft outline flexibly. The default settings for this template are set up to consider all items at the first level within the draft as chapters. Beneath that, all items will be considered sections, subsections and so on down to the paragraph environment. To adjust these settings, use the **Project****▸****?****Project Settings...** menu command, click on the “Section Types” pane, and then the “Default Types by Structure” tab. For example, if you wanted the top level of the Draft to use parts instead of chapters, you could select “Part” (which has already been set up for you) for the Level 1 assignments, and then adjust all greater levels of depth accordingly.

It is also worth noting that under default settings, if you omit a title for a binder item it will become "anonymous", not producing any title material and by extension opting out of being a part of the formal structure of the document as compiled. You could also manually assign an item to the “N/A” section type to cause it to be printed as-is, if you would prefer to give it an casual outlining title.

    ##  [Text]

Compiling

The “Front Matter” folder contains some example preamble of the sort that would *follow* the more rudimentary setup of the document's overall look and feel (that is delegated to the compile Format itself, making it possible to create your own target formats and switch document classes with a single click).

A “Back Matter” folder is also provided. Here you can add any back matter you wish to include before the final enddocument command. Nothing provided in this folder is necessary for a simple document to be typeset.  
  
An example structure has been provided for including an appendix. The folder itself is assigned to the "Appendix" Section Type, which is responsible for titling the section and then starting the formal appendix portion of the document. Within that folder the ordinary structural rules for the Draft folder will be adhered to, meaning you can construct complex appendices with sections, subsections and so forth, based on the hierarchy of the outline. If you do not need an appendix, you can remove this folder.

**To compile:**

Go to **File****▸****?****Compile...**

Next to “Compile for”, select “Plain Text (.txt)”.

Select the “LaTeX (Memoir Book)” compile format on the left.

Ensure that the “Add front matter” button is ticked under the contents list on the right. If you do not do this the resulting .tex file will not properly typeset until the example material has been copied to an area included in the compile (such as the top of the Draft folder). The “Add back matter” feature is optional.

Click on **Compile**, open the .tex file in your preferred editor and typeset the document.

**To create your own document class variation:**

Go to **File****▸****?****Compile...**

Right-click on the “LaTeX (Memoir Book)” compile format on the left, and select “Duplicate & Edit Format...”.

Rename the format at the top to match your intended edits.

Click on the Text Layout compile format preference pane.

Modify the prefix and suffix fields to include you preferred preamble and footer.

Since some of the LaTeX code used elsewhere in the settings may require adjustment to work outside of Memoir, you might need to audit the rest of the format's panes and content therein.

Click on **Save**.

    ##  [Text]

Sample Document

See the “” item in the Research folder for an example of a document that has been created using this template.

    ##  [Text]

Final Note

Scrivener project templates are flexible and are not intended to restrict you to a particular workflow. You can change, delete or move the files and folders contained in the template to suit how you work.

Like all templates in Scrivener, this project was originally created from the “Blank” template. We’ve simply added a few things to the binder and set everything up in ways that should be useful to authors working toward the LaTeX typesetting system. Everything you can do with this project, you could equally do by creating a “Blank” project and setting it up yourself.

You can create your own templates by setting up a skeletal project with the files, folders and settings you would like to use for new projects and using **File****▸****?****Save As Template**.

The rest of this guide will go over the numerous convenience features that have been added to this project.

  ## Usage Tips [Folder]

**Usage Tips**

This section will reference the various “features” this template provides out of the box, and in some cases, how to customise them or make your own, based upon the examples that have been given.

    ##  [Text]

Figures

To ease typing extensive syntax and managing your images by hand, or to simply grant yourself a nicer looking working document, you can make use of Scrivener’s native inline & linked image handling facilities. Graphics that you place into the editor will be exported as files and syntax will be inserted where the image was placed. For more formal figures there are two approaches you can take (you needn’t commit to either exclusively):

**In your text**: use the "Figure" paragraph style to have the graphic enclosed in a beginfigurecentering ... endfigure environment. You can optionally insert a caption below the image (still within the "Figure" paragraph style, so as to keep it within the environment), using the "Caption" character style. Follow the caption with a "Label" styled phrase that is a valid LaTeX style bookmark label, to give this figure a cross-reference you can refer to elsewhere. With an eye on the Format Bar, click into the three different elements below; image, caption and label, to see which styles are applied to them:



This is the caption. latexLogo

The example provided above would generate the following LaTeX syntax when compiled:

beginfigure[htbp]

centering

includegraphics[width=175pt,height=73pt]800px-LaTeX_logo.png

captionThis is the caption.

labellatexLogo

endfigure

**As a binder item**: if you prefer keeping figures in separate chunks in the binder, to facilitate searching and organisation, you can use the "Figure" Section Type setting, placing the graphic in the main editor (un-styled) and supplying the caption as custom metadata in the Inspector. (A  has been supplied for your convenience, and with an  item in the “Research” folder for your reference.) When using this technique figure labelling will be handled automatically for you, and captioning will be done in the inspector sidebar. To cross-reference to the figure in the text, refer to the following section on cross-referencing.

**Figure Customisation**

To modify how styled text and embedded images are converted to LaTeX code, visit the Markup and Styles compile format panes, when editing this project's "LaTeX (Memoir Book)" format. The Markup pane itself sets how graphics print themselves, which is the includegraphics line alone. The stylesheet settings for Figure, Caption and Label determine how the additional syntax is added around that.

To modify how images are converted to LaTeX when including them as binder items, visit the Section Layouts compile format pane and adjust the "Figure w/ Caption" layout's Prefix and Suffix tabs. As with styled text, the Markup pane is what is used to insert the graphic itself between the prefix and suffix.

    ##  [Text]

Equations

A few helpful packages for equation typesetting can be enabled with the  module, in the front matter folder. As with figures, equations can be added either to your binder as discrete items using dedicated Section Types, or into the text itself using Styles.

**In your text**

Use the "Equation" paragraph style to enclose the selected text within a beginfigure ... endfigure environment. Use the Label character style to supply a target label to the equation, if it should be numbered and be capable of being cross-referenced. Below is an example equation and label; note how the label character style is placed within the equation paragraph itself:

eq:custom-label

The example provided above would generate the following LaTeX syntax when compiled:

beginequation

Pleft(A=2middle|fracA^2B>4right) labeleq:custom-label

endequation

**As a binder item**

For those that prefer to keep major elements as separate binder items, this project has two document templates designed to faciliate the organisation of equations: "" and "". If you will require cross-referencing to refer to your equations, the first option is what you want. Numbered equations will be automatically labelled for you, with that label used by internal document links pointing to it, elsewhere in the Draft. Consequently, you should not supply your own label... code (or use the style) with the equation text itself. Additionally you should not use the Equation style itself, within an Equation document.

Inline equations, embedded in paragraph text, are easily formatted using raw LaTeX, but if you wish to keep your source material clean of syntax, a character style has been provided with this project. Simply paint the equation with the style, and it will add the markers for you.

    ##  [Text]

Tables

Scrivener cannot convert rich text tables into LaTeX code for you. You will need to provide the proper syntax yourself, however you can focus on the tabular data itself, leaving the boilerplate text to Scrivener by making use of the “Table” Section Type when applied to individual binder items. As with binder figures, you can supply a caption as metadata. A  has been provided as an example starting point. Working this way means you can make global adjustments to table appearance more easily, by editing the “Table w/ Caption” section layout, in the compile format settings.

    ##  [Text]

Itemised and Enumerated Lists

As with tables, Scrivener cannot convert its native listing environments to LaTeX. If you would prefer the convenience of these features, for simple lists you could make use of the “Itemized List” and “Enumerated List” styles, included in this template. You will note that simple hanging indent paragraph style will be used in the editor (along with a shaded background to indicate the type of list). When compiled with the provided format, these styles will be enclosed in the proper listing environment, and will prefix each line with the item code.

Given that basic capability, you could embellish the standard output the compiler produces with your own LaTeX code inside the styled range, and thus produce nested lists.

    ##  [Text]

Creating Your Own Environments

This template includes a variety of common environments already set up for you, but it is quite likely that you will need to create your own. There are two different approaches, both demonstrated in this template:

**Using Section Types**

Examples of this approach include the figure, table and equation setups. These are comprised of the following ingredients:

A “Section Type” that can be used to flag items in your binder as needing to be associated with a particular environment. If you look at the  document template, you will see that in the Metadata inspector pane, it has been assigned to the “Table” section type.

A “Section Layout” in the compile format settings is what does the heavy lifting. Edit the Format for this project by going to **File****▸****?****Compile...**, right-click on the project template's Format in the left sidebar to edit it, and in the “Section Layouts” pane, examine the “Table w/ Captions” layout. The prefix and suffix tabs for this layout are where the actual LaTeX code will be inserted around the text you type in the editor.

**Using Styles**

When environments are more likely to be found within a longer chunk of text, and would not be so useful broken out into their own sections, styles can also be used to insert a prefix and suffix around the text you type. The “Block Quote” or even “Emphasis” styles serve as examples of how this can be done. Visit the “Styles” pane in the compile format settings to examine how prefixes and suffixes are used with them.

    ##  [Text]

Cross-References

Whenever you create a document link to another document within the compiled output, Scrivener will use an internally generated bookmark label to create an autoref... link to that section. This works for the following Section Types automatically:

Part

Chapter

Section

Subsection

Subsubsection

Figure

Table

Equation

LaTeX will of course handle the rest of the job, in terms of numbering and correctly labelling the links. To adjust the behaviour of the code produced by Scrivener, modify the link prefix and suffix settings in the Markup compile format pane.

The provided settings work with the convention of ignoring the hyperlink formatting itself, save for the significance of where the link ends. Where it ends is where the autoref command will be inserted. Thus, the text of the link itself is meant to be printed as normal text that is readable by the user. You may link to however much or little text as you please.

**Cross-Reference Customisation**

The settings that maintain how links to internal documents will be marked up can be found in the Markup compile format option pane:

Open the compiler, and double-click on the “LaTeX (Memoir Book)” format in the left sidebar to edit it.

Select the Markup pane, and examine the **Internal link prefix and suffix** fields.

Given that you have control over what is inserted before and after the text that is hyperlinked in the editor, the flexibility for formatting links more expressively is available. You might for instance prefer underlined links for digital documents, without any auto-ref suffix.

    ##  [Text]

Typographic Punctuation

If you prefer to use LaTeX style punctuation yourself while writing, then Scrivener will not get in your way in doing so. However if you prefer to use its automatic typographic punctuation features (smart quotes, dashes and ellipses), the compile format has been configured to convert these forms of character punctuation to old style safe ASCII LaTeX methods. If you would rather work in a modern UTF-8 environment where typographic punctuation is acceptable, you could remove or disable the conversions from the Replacements compile format pane.

    ##  [Text]

Notation

If you opt to include inline annotations and comments in your output, they will be converted to marginal notes and highlights using LaTeX tools for doing so. You may also make use of Scrivener's footnote features, as they will be converted to syntax for you.

    ##  [Text]

Creating an Index

To an include an index of terms, you will first need to enable the  front matter module, along with the  itself in the back matter folder. To do so, visit these sections and enable the “Include in Compile” option in the inspector’s Metadata tab.

As for marking indexed terms themselves, you could of course do so by hand, but you can also make use of the Index Term and Index Key styles:

Use the “Index Term” style to mark the word or phrase, directly in the text, that you wish to add to the index. When compiled, a key based on the selected text will be generated for you.

Use the “Index Key” style if the phrase itself doesn’t appear in the text, or if you want direct control over the indexing key itself.

## Draft [DraftFolder]

  ## Chapter [Folder]

    ## Section [Text]

Testeando testeando

    ##  [Text]

Otro test otro test

# SUPER!

A ver si esto es de tu talla

## Front Matter [Folder]

  ## Preamble Modules [Folder]

    ## Equation Packages [Text]

% Math packages

usepackageamssymb

usepackageamsmath

usepackageamsthm

    ## Bibliography Packages [Text]

usepackagenatbib

    ## Indexing Packages [Text]

usepackagemakeidx

makeindex

    ## Proofing Packages [Text]

% CriticMarkup Support

% See document notes in the inspector for compile setting adjustments.

% Credit goes to Fletcher Penney, of MultiMarkdown, for these methods.

usepackagesoul

usepackagexargs

usepackagetodonotes

newcommandxcmnote[2][1=]linespread1.0todo[linecolor=red,backgroundcolor=red!25,bordercolor=red,#1]#2

% Use ul instead of underline since we are using soul

letunderlineul

% Use a wider margin for notes

% Use entire page

settrimmedsizestockheightstockwidth*

settrims0pt0pt

setlrmarginsandblock2.5cm5.5cm*

setulmarginsandblock3.5cm3.5cm*

checkandfixthelayout

  ## Begin Document [Text]

setcounterpage1

pagenumberingroman

title<$projecttitle>

author<$author>

date<$date>

begindocument

maketitle

clearpage

  ## Contents [Text]

tableofcontents

% listoffigures

% listoftables

  ## Foreward [Text]

Algo

  ## Begin Main Matter [Text]

newpage

setcounterpage1

pagenumberingarabic

mainmatter

## Back Matter [Folder]

  ##  [Text]

backmatter

  ## Appendices [Folder]

    ## Appendix [Text]

  ## Bibliography [Text]

bibliographystylespiebib

bibliographyname_of_bib_file

  ## Index [Text]

printindex

## Notes [Folder]

## Ideas [Folder]

## Research [ResearchFolder]

  ## Sample PDF [PDF]

  ## Sample PDF Material [Folder]

    ## Instructions for compiling sample PDF [Text]

Compilation Instructions

========================

If you would like to compile the sample PDF, then select the two folders below, and use the **Documents > Move To > Draft** menu command. The following adjustments should then be made to settings:

1. Enter **File > Compile...**

2. Enable the “Add back matter” feature, and select the “Back Matter” folder from the binder. Disable the two appendix related items, and enable the “Index” item.

    ## Folder Names as Chapter Titles [Folder]

      ##  [Text]

The goal of this template is make the composition and drafting phase of a LaTeX-based project as simple and elegant as possible. It is not meant to turn Scrivener, a rich text editor, into a full-blown LaTeX generator (you'll need to use MultiMarkdown or Pandoc for that; consider using the General Non-Fiction template if that is what you are looking for). Instead, this project contains a number of convenience features that could be applied to your work as you wish in an àà la carte fashion, to alleviate how much code you'd need to type in yourself. For example, if you find handling images to be cumbersome in LaTeX, you can safely leave that up to the template.

When compiled (File>Compile…), this project will generate a bare bones .tex document(using the flexible Memoir class) suitable for many types of work, . The template achieves this by using advanced features available to the Plain Text compile file type, including providing the default extension of . From that basic start you could build your own document design, or just as easily swap out the preamble components for something else entirely.

Foremost, this template caters to those that prefer, or perhaps are required, to compose directly in LaTeX itself, and will be approaching Scrivener primarily as an organisational and composition tool for constructing a longer .tex file from the many smaller snippets written into the Draft folder.

      ## How To Use This Template [Text]

Inside the Draft folder, create a new folder for each chapter and title each folder with the name of the chapter. (You do not need toand indeed shouldnttitle the folders Chapter One and so on, because chapter numbering will typically be taken care of automatically by LaTeX.) The first chapter folder has been created for you with the placeholder title Chapter.Create a new text document for each section inside the chapter folders. Sections will be numbered and titled by LaTeX using hierarchical numbering (1.1, 1.2, 1.3).Nest documents beneath documents to create numbered subsections, subsubsections and unnumbered paragraphs, at 3rd, 4th and 5th levels of hierarchy respectively.If you need to create the occasional unnumbered section, the Informal Chapter and Informal Section Types are there for you.If you dont require a foreword, move the Foreword document in the Front Matter folder to the Trash folder. Alternatively, rename it Preface, Abstract or Introduction as you require. Any other prefacing materials could be placed between the preamble and the main matter document.Notes and Ideas folders have been provided for your convenience, although you can replace them or create different top-level folders for your research materials if necessary (these are just regular folders that have had custom icons assigned to them using the Documents$>$Change Icon feature).

      ## Some Sample Content [Text]

Su galph velar; ewayf, xu anu srung gen wynlarce frimba fli erc kurnap furng, xi kurnap, clum furng erk xi ti harle helk irpsa xu quolt. Gronk teng furng xu obrikt ju---ik clum rintax xu prinquis la. Prinquis zorl ewayf; urfa dri, vo xi, ju fli zeuhl obrikt re gronk teng morvit wex irpsa. Ma nix ewayf lamax. La furng nix brul ewayf sernag korsa yem tharn. Ju teng. Arul vusp er rintax athran, re sernag... ma prinquis. Urfa furng athran jince, gra nix tolaspa la delm; gronk dwint epp yiphras.

        ## Figures [Text]

This is the caption. ExampleImage

To modify how styled text is converted to LaTeX code, visit the Markup and Styles compile format panes, when editing this project's "LaTeX (Memoir Book)" format. The Markup pane itself sets how graphics print themselves, which is the textbackslashincludegraphics line alone. The stylesheet settings for Figure, Caption and Label determine how the additional syntax is added around that.

        ## Equations [Text]

We can also insert equations as discrete items in the binder, making them easy to reference and look up by type. It also means we can refer to them in our writings with the use of simple . In a similar fashion, we can also refer to .

          ## Equations [Text]

x_2 = frac5 - sqrt25 - 4 times 62 = 2

          ##  [Text]

When you need to print to an equation in the body text beta = (beta_1,beta_2 ldots beta_n) you can use the Inline Equation style which will insert the ( and ) codes for you.

Equations can also be inserted unnumbered:

          ## Unnumbered Equations [Text]

Pleft(A=2middle|fracA^2B>4right)

        ##  [Text]

Athran teng ti anu su, thung... ti quolt ma xi rintax yem urfa obrikt frimba flim furng. Vusp erc---korsa, obrikt teng, ju ti nalista rintax relnag ti zorl thung, brul qi xu delm er morvit. Ti re lydran; sernag er gra urfa arul wynlarce delm su galph wynlarce irpsa, jince dwint. Berot sernag groum irpsa epp jince. Sernag dri; er gen relnag ik rintax er, erc su dri, teng groum dwint ju groum relnag. Irpsa lamax erc velar brul flim dwint kurnap irpsa wynlarce clum su menardis yiphras; zorl morvit delm korsa zeuhl fli srung. Wynlarce delm; gen nix, teng su ti ma, dwint su, re furng vo tolaspa obrikt groum lydran xi.

        ##  [Text]

Here is some verse.We do not need much more than a few lines.And then we will be done.

        ##  [Text]

Erc xi---helk, korsa ma velar vusp gen berot? Arul galph; flim groum su teng twock er velar arul, urfa ma tharn. Quolt cree pank menardis, wynlarce ti sernag su tharn lamax? Tolaspa zorl pank... vo prinquis relnag rintax tolaspa er gronk. Ozlint gibberish delm gra; dwint er zorl fli prinquis xi gen qi srung, erc xu whik prinquis irpsa re. Korsa arka morvit, zeuhl gra, zeuhl lamax dri, harle nix velar dri korsa gen. Ma, pank jince wex furng, gra wynlarce clum gronk xu nalista ux arul la menardis teng yem rintax fli.Su xi er relnag tolaspa arka---jince erc twock morvit, su ux ozlint korsa ma. Thung epp, yem; yiphras berot la ma vusp su brul prinquis rhull pank vusp relnag erc ma rintax? Qi ik pank delm zorl lamax tolaspa lamax kurnap yem whik ju. Fli morvit---dri su srung gronk cree furng xi. Obrikt gen wex, arul ozlint, ma korsa rintax srung, epp ux, er pank. Lamax ozlint arul epp ju; qi ik ti harle ik. Lamax su ozlint ju ma yem rhull brul rintax re athran flim gra; yiphras rhull wex epp. Er ik epp erk zorl ewayf, ozlint re; wynlarce twock clum, irpsa ewayf er gronk. Su galph velar; ewayf, xu anu srung gen wynlarce frimba fli erc kurnap furng, xi kurnap, clum furng erk xi ti harle helk irpsa xu quolt. Gronk teng furng xu obrikt ju---ik clum rintax xu prinquis la. Prinquis zorl ewayf; urfa dri, vo xi, ju fli zeuhl obrikt re gronk teng morvit wex irpsa. Ma nix ewayf lamax. La furng nix brul ewayf sernag korsa yem tharn. Ju teng. Arul vusp er rintax athran, re sernag... ma prinquis. Urfa furng athran jince, gra nix tolaspa la delm; gronk dwint epp yiphras.

    ## Fli Zeuhl Obrikt [Folder]

      ##  [Text]

on a new top level folder like this. With LaTeX, the chapter numbering and formatting is handled automatically by the typesetting engine. Scrivener is only responsible for assembling the LaTeX code that will be used to generate these structures.

If you need parts,  don't need chapters or just  prefer to use single-file chapters, the automatic assignment of structure to the outline can be adjusted in project settings, under the Section Types pane.

      ## Jince erc Twock su Morvit Ux [Text]

Whik gronk; thung epp rintax whik jince dwint srung sernag nix la quolt sernag brul jince. Twock, quolt whik tharn dri cree gen... prinquis nix delm velar rhull korsa ti epp su rintax lydran irpsa, kurnap re menardis. Ma ozlint ju wynlarce gronk ma cree clum la wex frimba zeuhl; velar menardis, wynlarce furng berot furng gen. Thung er wynlarce wex tolaspa, srung morvit galph. Gen athran morvit... korsa, morvit menardis kurnap rintax velar teng srung vo frimba. Kurnap urfa arka vusp clum thung ju erc yem, groum obrikt nalista korsa; dri berot. Groum galph; ik, morvit ti gronk zeuhl erc nix. Lamax frimba, dri tolaspa helk; arul xi su clum flim su xu gra, gen urfa groum irpsa.beginitemize

item<u></u><u>Using Section Types</u>: examples of this approach include the figure, table and equation setups. These are comprised of the following ingredients:beginenumerate

item A Section Type that can be used to flag items in your binder as needing to be associated with a particular environment. If you look at the Table document template, you will see that in the Metadata inspector pane, it has been assigned to the Table section type.item A Section Layout in the compile format settings is what does the heavy lifting. Edit the Format for this project by going to File▸Compile, right-click on the project template's Format in the left sidebar to edit it, and in the Section Layouts pane, examine the Table w/ Captions layout. The prefix and suffix tabs for this layout are where the actual LaTeX code will be inserted around the text you type in the editor.

endenumerateitem<u></u><u>Using Styles</u>: when environments are more likely to be found within a longer chunk of text, and would not be so useful broken out into their own sections, styles can also be used to insert a prefix and suffix around the text you type. The Block Quote or even Emphasis styles serve as examples of how this can be done. Visit the Styles pane in the compile format settings to examine how prefixes and suffixes are used with them.

enditemizeUrfa erc prinquis; tharn yem arka, vusp xu erc. Fli xi menardis arka... ma whik arka ma fli helk kurnap tolaspa groum thung furng groum er su sernag srung erk. Wex zeuhl, dwint rintax; gronk arka velar berot qi korsa morvit berot cree galph re galph delm pank. Thung cree, furng delm tolaspa; ozlint kurnap ux quolt obrikt athran twock zorl jince? Re groum; thung su flim kurnap su vo quolt, wex er zorl gen xu ti re. Wynlarce, ti prinquis ux lamax gen wex, wynlarce er la erk lamax rhull? Delm vo; berot nix erc twock wynlarce gronk ju? Yem groum whik erk galph urfa epp; kurnap nalista brul, zeuhl vo. Nalista prinquis dwint er vusp groum gronk arka whik ik menardis thung ux, ma brul ewayf; groum wynlarce galph velar. Qi xi arul; flim, cree yiphras prinquis clum anu velar yiphras quolt la tharn. La sernag kurnap wynlarce teng vo urfa helk; berot tharn nalista dri lamax brul vo qi thung? Galph wex ma epp, twock relnag berot. Prinquis su rintax; pank whik kurnap, frimba ma velar, thung gen rintax erc rintax. Ju ti erk gronk ewayf ux, nix prinquis frimba. Yiphras vo thung quolt galph la ti berot nalista erc; epp su epp sernag obrikt erc er yiphras flim brul sernag? Obrikt, whik teng srung nix prinquis xi brul lydran re urfa... groum ti, er harle dri quolt menardis groum qi. Harle gra harle irpsa, la zeuhl. Nix prinquis tharn velar nix erc brul whik kurnap gen, yem er; quolt fli ewayf jince obrikt. Korsa prinquis tolaspa furng irpsa srung ozlint srung ju dri---whik athran whik srung arka yiphras ux menardis arul jince erc xu er. Athran teng ti anu su, thung... ti quolt ma xi rintax yem urfa obrikt frimba flim furng. Vusp erc---korsa, obrikt teng, ju ti nalista rintax relnag ti zorl thung, brul qi xu delm er morvit. Ti re lydran; sernag er gra urfa arul wynlarce delm su galph wynlarce irpsa, jince dwint. Berot sernag groum irpsa epp jince. Sernag dri; er gen relnag ik rintax er, erc su dri, teng groum dwint ju groum relnag. Irpsa lamax erc velar brul flim dwint kurnap irpsa wynlarce clum su menardis yiphras; zorl morvit delm korsa zeuhl fli srung. Wynlarce delm; gen nix, teng su ti ma, dwint su, re furng vo tolaspa obrikt groum lydran xi. Erc xi---helk, korsa ma velar vusp gen berot? Arul galph; flim groum su teng twock er velar arul, urfa ma tharn. Quolt cree pank menardis, wynlarce ti sernag su tharn lamax? Tolaspa zorl pank... vo prinquis relnag rintax tolaspa er gronk. Ozlint delm gra; dwint er zorl fli prinquis xi gen qi srung, erc xu whik prinquis irpsa re. Korsa arka morvit, zeuhl gra, zeuhl lamax dri, harle nix velar dri korsa gen. Ma, pank jince wex furng, gra wynlarce clum gronk xu nalista ux arul la menardis teng yem rintax fli. Su xi er relnag tolaspa arka---jince erc twock morvit, su ux ozlint korsa ma. Thung epp, yem; yiphras berot la ma vusp su brul prinquis rhull pank vusp relnag erc ma rintax? Qi ik pank delm zorl lamax tolaspa lamax kurnap yem whik ju. Fli morvit---dri su srung gronk cree furng xi. Obrikt gen wex, arul ozlint, ma korsa rintax srung, epp ux, er pank. Lamax ozlint arul epp ju; qi ik ti harle ik. Lamax su ozlint ju ma yem rhull brul rintax re athran flim gra; yiphras rhull wex epp. Er ik epp erk zorl ewayf, ozlint re; wynlarce twock clum, irpsa ewayf er gronk. Su galph velar; ewayf, xu anu srung gen wynlarce frimba fli erc kurnap furng, xi kurnap, clum furng erk xi ti harle helk irpsa xu quolt. Gronk teng furng xu obrikt ju---ik clum rintax xu prinquis la. Prinquis zorl ewayf; urfa dri, vo xi, ju fli zeuhl obrikt re gronk teng morvit wex irpsa. Ma nix ewayf lamax. La furng nix brul ewayf sernag korsa yem tharn. Ju teng. Arul vusp er rintax athran, re sernag... ma prinquis. Urfa furng athran jince, gra nix tolaspa la delm; gronk dwint epp yiphras.

        ## A subsection [Text]

Er xu kurnap velar ik ti quolt ozlint. Frimba kurnap ux brul furng; ju tharn helk yem su groum frimba. Srung gronk erk er whik prinquis teng er galph; berot arul lamax erc gra. Relnag ti. Kurnap sernag dwint frimba wynlarce---urfa harle pank su berot arul, zorl harle vo pank. Whik erc---erk jince korsa ma dwint wex nix athran thung ma, twock sernag urfa gen tharn. Ti twock velar, athran yiphras---obrikt clum rhull nalista ozlint ik re korsa. Brul kurnap er sernag obrikt delm erc nix, anu er su tolaspa gen tolaspa nix furng whik la furng korsa? Twock gra harle, yem nix; clum, erk berot gronk rintax gen teng ju, prinquis ma xu, ozlint nalista? Sernag rintax re gen velar tolaspa xi prinquis; su jince erk urfa? Thung clum tolaspa thung ik erk; dwint gra, nalista thung, rhull ux. Zorl tharn srung kurnap velar xu wex er dwint gra vo zorl; irpsa cree la vo lydran, dwint jince prinquis teng prinquis. Zorl dri erk nix, ux tolaspa whik, arul er, vo xi anu zeuhl quolt; galph helk korsa flim ma vo. Su er berot gen yem srung whik ux arul dri quolt whik. Helk groum clum; helk yem berot er relnag harle, teng obrikt helk athran quolt wex jince quolt helk harle, relnag lydran furng er ik rintax tharn la. Frimba lydran er; nalista whik lamax menardis ik sernag, delm jince helk morvit gen, erk kurnap flim erc ik athran prinquis pank yiphras erk, vo tolaspa. Arul dwint relnag xi ma, dri dwint harle. Kurnap menardis galph ik gronk sernag gronk delm morvit vusp yiphras jince ju. Erk flim anu rhull anu, lydran wex furng morvit delm. Pank nalista, whik nix vusp obrikt dri thung nalista morvit zeuhl, kurnap berot qi prinquis berot. Twock erk erc er... morvit galph ewayf frimba morvit sernag er prinquis? Vo prinquis zeuhl, anu zeuhl la arka erc epp, zeuhl re relnag wynlarce su urfa. Su, er velar cree brul yem, gronk helk anu gen athran yiphras ti athran harle ju quolt gen menardis urfa epp ewayf groum. Dwint qi athran sernag anu kurnap arul. Er fli yiphras sernag vo kurnap dwint vusp obrikt rhull gen. Helk yiphras er arul wynlarce korsa arul... korsa tolaspa anu wynlarce lamax srung. Ux ik xu rintax nix. Sernag teng, yiphras erk tolaspa---gen ux velar ik irpsa menardis arul ux obrikt yem zeuhl. Lamax su dri er anu; ewayf er jince, wex er arul flim erk fli. Relnag jince wynlarce jince prinquis. Tharn jince vusp er... gra, korsa yem. Dri la berot ozlint; clum twock, anu rhull dri helk, vusp zorl rintax obrikt wex. Twock lydran arka harle re nalista groum quolt; menardis ma brul qi pank berot dri morvit, erc kurnap delm.

  ## Example figure [Text]

  ## Templates [Folder]

    ## Figure [Text]

*<INSERT IMAGE HERE>*

    ## Table [Text]

begintabularll

textbfColumn 1 & textbfColumn 2

midline

a & b 

endtabular

    ## LaTeX Control File [Text]

% This file type can be used to insert custom commands, preamble settings, or whatever material should be printed directly rather than treated as part of the structure of the work.

    ## Equation [Text]

    ## Unnumbered Equation [Text]

